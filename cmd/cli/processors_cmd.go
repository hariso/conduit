package cli

import (
	"context"
	"fmt"

	"github.com/conduitio/conduit/pkg/conduit"
	"github.com/conduitio/conduit/pkg/foundation/database/inmemory"
	"github.com/conduitio/conduit/pkg/pipeline"
	"github.com/conduitio/conduit/pkg/provisioning"
	"github.com/spf13/cobra"
)

func cmdProcessors() *cobra.Command {
	cfg := conduit.DefaultConfig()
	ctx := context.Background()

	processorsCmd := &cobra.Command{
		Use:   "processors",
		Short: "Manage Conduit Processors",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	// 1. Create a new command to list pipelines
	// 2. Add that command as `pipelines ls` to the pipelineCmd

	processorsListCmd := &cobra.Command{
		Use:     "ls",
		Aliases: []string{"list"},
		Short:   "List all processors",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: Check if I'm running the service correctly

			pipelinePath, _ := cmd.Flags().GetString("path")

			logger := conduit.NewLogger(cfg.Log.Level, cfg.Log.Format)
			db := &inmemory.DB{}
			pipelineService := pipeline.NewService(logger, db)

			provisionService := provisioning.NewService(db, logger, pipelineService, nil, nil, nil, pipelinePath)

			r := &conduit.Runtime{
				Config:           cfg,
				DB:               db,
				Ready:            make(chan struct{}),
				Logger:           logger,
				ProvisionService: provisionService,
				PipelineService:  pipelineService,
			}
			r.ProvisionService.Init(ctx)

			// TODO: Init without provision

			r.PipelineService.Init(ctx)
			r.PipelineService.Run(ctx, nil, nil, nil)

			instances := r.PipelineService.List(ctx)
			for _, instance := range instances {
				fmt.Println(instance)
			}

			fmt.Println("Flag Value:", pipelinePath)
		},
	}
	processorsListCmd.PersistentFlags().StringVar(&cfg.Pipelines.Path, "path", cfg.Pipelines.Path, "path to the directory that has the yaml pipeline configuration files, or a single pipeline configuration file")

	processorsCmd.AddCommand(processorsListCmd)

	return processorsCmd
}
