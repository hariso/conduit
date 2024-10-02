// Copyright © 2024 Meroxa, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package lifecycle contains the logic to manage the lifecycle of pipelines.
// It is responsible for starting, stopping and managing pipelines.
package lifecycle

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/conduitio/conduit-commons/csync"
	"github.com/conduitio/conduit/pkg/connector"
	"github.com/conduitio/conduit/pkg/foundation/cerrors"
	"github.com/conduitio/conduit/pkg/foundation/log"
	"github.com/conduitio/conduit/pkg/foundation/metrics/measure"
	"github.com/conduitio/conduit/pkg/lifecycle/funnel"
	"github.com/conduitio/conduit/pkg/lifecycle/stream"
	"github.com/conduitio/conduit/pkg/pipeline"
	connectorPlugin "github.com/conduitio/conduit/pkg/plugin/connector"
	"github.com/conduitio/conduit/pkg/processor"
	"github.com/jpillora/backoff"
	"gopkg.in/tomb.v2"
)

type FailureEvent struct {
	// ID is the ID of the pipeline which failed.
	ID    string
	Error error
}

type FailureHandler func(FailureEvent)

// Service manages pipelines.
type Service struct {
	logger log.CtxLogger

	backoffCfg *backoff.Backoff

	pipelines  PipelineService
	connectors ConnectorService

	processors       ProcessorService
	connectorPlugins ConnectorPluginService

	handlers         []FailureHandler
	runningPipelines *csync.Map[string, *runnablePipeline]
}

// NewService initializes and returns a lifecycle.Service.
func NewService(
	logger log.CtxLogger,
	backoffCfg *backoff.Backoff,
	connectors ConnectorService,
	processors ProcessorService,
	connectorPlugins ConnectorPluginService,
	pipelines PipelineService,
) *Service {
	return &Service{
		logger:           logger.WithComponent("lifecycle.Service"),
		backoffCfg:       backoffCfg,
		connectors:       connectors,
		processors:       processors,
		connectorPlugins: connectorPlugins,
		pipelines:        pipelines,
		runningPipelines: csync.NewMap[string, *runnablePipeline](),
	}
}

type runnablePipeline struct {
	pipeline *pipeline.Instance
	w        *funnel.Worker
	t        *tomb.Tomb
}

// ConnectorService can fetch and create a connector instance.
type ConnectorService interface {
	Get(ctx context.Context, id string) (*connector.Instance, error)
	Create(ctx context.Context, id string, t connector.Type, plugin string, pipelineID string, cfg connector.Config, p connector.ProvisionType) (*connector.Instance, error)
}

// ProcessorService can fetch a processor instance and make a runnable processor from it.
type ProcessorService interface {
	Get(ctx context.Context, id string) (*processor.Instance, error)
	MakeRunnableProcessor(ctx context.Context, i *processor.Instance) (*processor.RunnableProcessor, error)
}

// ConnectorPluginService can create a connector plugin dispenser.
type ConnectorPluginService interface {
	NewDispenser(logger log.CtxLogger, name string, connectorID string) (connectorPlugin.Dispenser, error)
}

// PipelineService can fetch, list and update the status of a pipeline instance.
type PipelineService interface {
	Get(ctx context.Context, pipelineID string) (*pipeline.Instance, error)
	List(ctx context.Context) map[string]*pipeline.Instance
	UpdateStatus(ctx context.Context, pipelineID string, status pipeline.Status, errMsg string) error
}

// OnFailure registers a handler for a lifecycle.FailureEvent.
// Only errors which happen after a pipeline has been started
// are being sent.
func (s *Service) OnFailure(handler FailureHandler) {
	s.handlers = append(s.handlers, handler)
}

// Init starts all pipelines that have the StatusSystemStopped.
func (s *Service) Init(
	ctx context.Context,
) error {
	var errs []error
	s.logger.Debug(ctx).Msg("initializing pipelines statuses")

	instances := s.pipelines.List(ctx)
	for _, instance := range instances {
		if instance.GetStatus() == pipeline.StatusSystemStopped {
			err := s.Start(ctx, instance.ID)
			if err != nil {
				// try to start remaining pipelines and gather errors
				errs = append(errs, err)
			}
		}
	}

	return cerrors.Join(errs...)
}

// Start builds and starts a pipeline with the given ID.
// If the pipeline is already running, Start returns ErrPipelineRunning.
func (s *Service) Start(
	ctx context.Context,
	pipelineID string,
) error {
	pl, err := s.pipelines.Get(ctx, pipelineID)
	if err != nil {
		return err
	}

	if pl.GetStatus() == pipeline.StatusRunning {
		return cerrors.Errorf("can't start pipeline %s: %w", pl.ID, pipeline.ErrPipelineRunning)
	}

	s.logger.Debug(ctx).Str(log.PipelineIDField, pl.ID).Msg("starting pipeline")
	s.logger.Trace(ctx).Str(log.PipelineIDField, pl.ID).Msg("building tasks")

	rp, err := s.buildRunnablePipeline(ctx, pl)
	if err != nil {
		return cerrors.Errorf("could not build tasks for pipeline %s: %w", pl.ID, err)
	}

	s.logger.Trace(ctx).Str(log.PipelineIDField, pl.ID).Msg("running pipeline")

	if err := s.runPipeline(rp); err != nil {
		return cerrors.Errorf("failed to run pipeline %s: %w", pl.ID, err)
	}
	s.logger.Info(ctx).Str(log.PipelineIDField, pl.ID).Msg("pipeline started")

	s.runningPipelines.Set(pl.ID, rp)

	return nil
}

// Stop will attempt to gracefully stop a given pipeline by calling each worker's
// Stop method. If the force flag is set to true, the pipeline will be stopped
// forcefully by cancelling the context.
func (s *Service) Stop(ctx context.Context, pipelineID string, force bool) error {
	rp, ok := s.runningPipelines.Get(pipelineID)

	if !ok {
		return cerrors.Errorf("pipeline %s is not running: %w", pipelineID, pipeline.ErrPipelineNotRunning)
	}

	if rp.pipeline.GetStatus() != pipeline.StatusRunning && rp.pipeline.GetStatus() != pipeline.StatusRecovering {
		return cerrors.Errorf("can't stop pipeline with status %q: %w", rp.pipeline.GetStatus(), pipeline.ErrPipelineNotRunning)
	}

	switch force {
	case false:
		s.logger.Info(ctx).
			Str(log.PipelineIDField, rp.pipeline.ID).
			Any(log.PipelineStatusField, rp.pipeline.GetStatus()).
			Msg("gracefully stopping pipeline")
		rp.w.Stop()
	case true:
		s.logger.Info(ctx).
			Str(log.PipelineIDField, rp.pipeline.ID).
			Any(log.PipelineStatusField, rp.pipeline.GetStatus()).
			Msg("force stopping pipeline")
		rp.t.Kill(pipeline.ErrForceStop)
	}
	return nil
}

func (s *Service) stopGraceful(ctx context.Context, rp *runnablePipeline) {
}

// StopAll will ask all the running pipelines to stop gracefully
// (i.e. that existing messages get processed but not new messages get produced).
func (s *Service) StopAll(ctx context.Context, reason error) {
	if s.runningPipelines.Len() == 0 {
		return
	}

	for _, rp := range s.runningPipelines.All() {
		if rp.pipeline.GetStatus() != pipeline.StatusRunning && rp.pipeline.GetStatus() != pipeline.StatusRecovering {
			continue
		}

		s.logger.Info(ctx).
			Str(log.PipelineIDField, rp.pipeline.ID).
			Any(log.PipelineStatusField, rp.pipeline.GetStatus()).
			Msg("gracefully stopping pipeline")
		rp.w.Stop()
	}

	// Wait for the pipelines to stop
	const (
		interval = time.Second * 10
		count    = 6
	)

	var err error
	for i := count; i > 0; i-- {
		if i == 1 {
			// on last try, stop forcefully
			running := s.runningPipelines.Len()
			s.logger.Warn(ctx).Msgf("Stopping (%d) pipelines forcefully", running)

			for _, rp := range s.runningPipelines.All() {
				rp.t.Kill(pipeline.ErrForceStop)
			}
		}

		s.logger.Info(ctx).Msgf("Waiting for %d pipelines to stop (time left: %s)", s.runningPipelines.Len(), time.Duration(i)*interval)
		err = s.Wait(interval)
		if err == nil {
			break
		}
	}

	if err != nil {
		s.logger.Warn(ctx).Msgf("Some pipelines failed to stop (%d)", s.runningPipelines.Len())
		return
	}

	s.logger.Info(ctx).Msg("All pipelines stopped")
}

// Wait blocks until all pipelines are stopped or until the timeout is reached.
// Returns:
//
// (1) nil if all the pipelines are gracefully stopped,
//
// (2) an error, if the pipelines could not have been gracefully stopped,
//
// (3) ErrTimeout if the pipelines were not stopped within the given timeout.
func (s *Service) Wait(timeout time.Duration) error {
	gracefullyStopped := make(chan struct{})
	var err error
	go func() {
		defer close(gracefullyStopped)
		err = s.waitInternal()
	}()

	select {
	case <-gracefullyStopped:
		return err
	case <-time.After(timeout):
		return pipeline.ErrTimeout
	}
}

// waitInternal blocks until all pipelines are stopped and returns an error if any of
// the pipelines failed to stop gracefully.
func (s *Service) waitInternal() error {
	var errs []error

	// copy pipelines to keep the map unlocked while we iterate it
	pipelines := s.runningPipelines.Copy()

	for _, rp := range pipelines.All() {
		if rp.t == nil {
			continue
		}
		err := rp.t.Wait()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return cerrors.Join(errs...)
}

// WaitPipeline blocks until the pipeline with the given ID is stopped.
func (s *Service) WaitPipeline(id string) error {
	p, ok := s.runningPipelines.Get(id)
	if !ok || p.t == nil {
		return nil
	}
	return p.t.Wait()
}

// buildRunnablePipeline will build and connect all tasks configured in the pipeline.
func (s *Service) buildRunnablePipeline(
	ctx context.Context,
	pl *pipeline.Instance,
) (*runnablePipeline, error) {
	pipelineLogger := s.logger
	pipelineLogger.Logger = pipelineLogger.Logger.With().Str(log.PipelineIDField, pl.ID).Logger()

	srcTasks, err := s.buildSourceTasks(ctx, pl, pipelineLogger)
	if err != nil {
		return nil, cerrors.Errorf("failed to build source tasks: %w", err)
	}
	if len(srcTasks) == 0 {
		return nil, cerrors.New("can't build pipeline without any source connectors")
	}
	if len(srcTasks) > 1 {
		// TODO(multi-connector): remove check
		return nil, cerrors.New("pipelines with multiple source connectors currently not supported, please disable the experimental feature flag")
	}

	destTasks, err := s.buildDestinationTasks(ctx, pl, pipelineLogger)
	if err != nil {
		return nil, cerrors.Errorf("failed to build destination tasks: %w", err)
	}
	if len(destTasks) == 0 {
		return nil, cerrors.New("can't build pipeline without any destination connectors")
	}
	if len(destTasks) > 1 {
		// TODO(multi-connector): remove check
		return nil, cerrors.New("pipelines with multiple destination connectors currently not supported, please disable the experimental feature flag")
	}

	dlq, err := s.buildDLQ(ctx, pl, pipelineLogger)
	if err != nil {
		return nil, cerrors.Errorf("failed to build DLQ: %w", err)
	}

	worker, err := funnel.NewWorker(
		[]funnel.Task{srcTasks[0], destTasks[0]},
		[][]int{{1}, {}},
		dlq,
		pipelineLogger,
		measure.PipelineExecutionDurationTimer.WithValues(pl.Config.Name),
	)
	if err != nil {
		return nil, cerrors.Errorf("failed to create worker: %w", err)
	}
	return &runnablePipeline{
		pipeline: pl,
		w:        worker,
	}, nil
}

func (s *Service) buildProcessorNodes(
	ctx context.Context,
	pl *pipeline.Instance,
	processorIDs []string,
	first stream.PubNode,
	last stream.SubNode,
) ([]stream.Node, error) {
	var nodes []stream.Node

	prev := first
	for _, procID := range processorIDs {
		instance, err := s.processors.Get(ctx, procID)
		if err != nil {
			return nil, cerrors.Errorf("could not fetch processor: %w", err)
		}

		runnableProc, err := s.processors.MakeRunnableProcessor(ctx, instance)
		if err != nil {
			return nil, err
		}

		var node stream.PubSubNode
		if instance.Config.Workers > 1 {
			node = s.buildParallelProcessorNode(pl, runnableProc)
		} else {
			node = s.buildProcessorNode(pl, runnableProc)
		}

		node.Sub(prev.Pub())
		prev = node

		nodes = append(nodes, node)
	}

	last.Sub(prev.Pub())
	return nodes, nil
}

func (s *Service) buildParallelProcessorNode(
	pl *pipeline.Instance,
	proc *processor.RunnableProcessor,
) *stream.ParallelNode {
	return &stream.ParallelNode{
		Name: proc.ID + "-parallel",
		NewNode: func(i int) stream.PubSubNode {
			n := s.buildProcessorNode(pl, proc)
			n.Name = n.Name + "-" + strconv.Itoa(i) // add suffix to name
			return n
		},
		Workers: proc.Config.Workers,
	}
}

func (s *Service) buildProcessorNode(
	pl *pipeline.Instance,
	proc *processor.RunnableProcessor,
) *stream.ProcessorNode {
	return &stream.ProcessorNode{
		Name:           proc.ID,
		Processor:      proc,
		ProcessorTimer: measure.ProcessorExecutionDurationTimer.WithValues(pl.Config.Name, proc.Plugin),
	}
}

func (s *Service) buildSourceTasks(
	ctx context.Context,
	pl *pipeline.Instance,
	logger log.CtxLogger,
) ([]funnel.Task, error) {
	var tasks []funnel.Task

	for _, connID := range pl.ConnectorIDs {
		instance, err := s.connectors.Get(ctx, connID)
		if err != nil {
			return nil, cerrors.Errorf("could not fetch connector: %w", err)
		}

		if instance.Type != connector.TypeSource {
			continue // skip any connector that's not a source
		}

		src, err := instance.Connector(ctx, s.connectorPlugins)
		if err != nil {
			return nil, err
		}

		tasks = append(
			tasks,
			funnel.NewSourceTask(
				instance.ID,
				src.(*connector.Source),
				logger,
			),
		)
	}

	return tasks, nil
}

func (s *Service) buildDestinationTasks(
	ctx context.Context,
	pl *pipeline.Instance,
	logger log.CtxLogger,
) ([]funnel.Task, error) {
	var tasks []funnel.Task

	for _, connID := range pl.ConnectorIDs {
		instance, err := s.connectors.Get(ctx, connID)
		if err != nil {
			return nil, cerrors.Errorf("could not fetch connector: %w", err)
		}

		if instance.Type != connector.TypeDestination {
			continue // skip any connector that's not a destination
		}

		dest, err := instance.Connector(ctx, s.connectorPlugins)
		if err != nil {
			return nil, err
		}

		tasks = append(
			tasks,
			funnel.NewDestinationTask(
				instance.ID,
				dest.(*connector.Destination),
				logger,
				measure.ConnectorExecutionDurationTimer.WithValues(
					pl.Config.Name,
					instance.Plugin,
					strings.ToLower(instance.Type.String()),
				),
				measure.ConnectorBytesHistogram.WithValues(
					pl.Config.Name,
					instance.Plugin,
					strings.ToLower(instance.Type.String()),
				),
			),
		)
	}

	return tasks, nil
}

func (s *Service) buildDLQ(
	ctx context.Context,
	pl *pipeline.Instance,
	logger log.CtxLogger,
) (*funnel.DLQ, error) {
	conn, err := s.connectors.Create(
		ctx,
		pl.ID+"-dlq",
		connector.TypeDestination,
		pl.DLQ.Plugin,
		pl.ID,
		connector.Config{
			Name:     pl.ID + "-dlq",
			Settings: pl.DLQ.Settings,
		},
		connector.ProvisionTypeDLQ, // the provision type ensures the connector won't be persisted
	)
	if err != nil {
		return nil, cerrors.Errorf("failed to create DLQ destination: %w", err)
	}

	dest, err := conn.Connector(ctx, s.connectorPlugins)
	if err != nil {
		return nil, err
	}

	return funnel.NewDLQ(
		"dlq",
		dest.(*connector.Destination),
		logger,
		measure.DLQExecutionDurationTimer.WithValues(pl.Config.Name, conn.Plugin),
		measure.DLQBytesHistogram.WithValues(pl.Config.Name, conn.Plugin),
		pl.DLQ.WindowSize,
		pl.DLQ.WindowNackThreshold,
	), nil
}

func (s *Service) runPipeline(rp *runnablePipeline) error {
	if rp.t != nil && rp.t.Alive() {
		return pipeline.ErrPipelineRunning
	}

	// the tomb is responsible for running goroutines related to the pipeline
	rp.t = &tomb.Tomb{}
	ctx := rp.t.Context(nil)

	err := rp.w.Open(ctx)
	if err != nil {
		return cerrors.Errorf("failed to open worker: %w", err)
	}

	rp.t.Go(func() error {
		doErr := rp.w.Do(ctx)
		s.logger.Err(ctx, doErr).Str(log.PipelineIDField, rp.pipeline.ID).Msg("Pipeline worker stopped")

		closeErr := rp.w.Close(context.Background())
		err := cerrors.Join(doErr, closeErr)
		if err != nil {
			return cerrors.Errorf("worker stopped with error: %w", err)
		}

		return nil
	})

	return s.pipelines.UpdateStatus(ctx, rp.pipeline.ID, pipeline.StatusRunning, "")
}

// notify notifies all registered FailureHandlers about an error.
func (s *Service) notify(pipelineID string, err error) {
	if err == nil {
		return
	}
	e := FailureEvent{
		ID:    pipelineID,
		Error: err,
	}
	for _, handler := range s.handlers {
		handler(e)
	}
}
