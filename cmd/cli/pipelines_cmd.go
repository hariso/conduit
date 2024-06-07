package cli

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/conduitio/conduit/pkg/conduit"
	"github.com/spf13/cobra"
)

func isNotDSStore(path string, d fs.DirEntry) bool {
	return filepath.Base(path) != ".DS_Store"
}

func openFile(filename string) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "nano"
	}
	editorCmd := exec.Command(editor, filename)
	if err := editorCmd.Run(); err != nil {
		fmt.Println(err)
	}
}

func listDirectory(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	entries := make([]fs.DirEntry, 0, len(files))
	for _, file := range files {
		if !isNotDSStore(file.Name(), file) {
			continue
		}
		entries = append(entries, file)
	}

	for _, file := range entries {
		fmt.Println(file.Name())
	}

	return nil
}

func cmdPipelines() *cobra.Command {
	cfg := conduit.DefaultConfig()
	//ctx := context.Background()

	pipelineCmd := &cobra.Command{
		Use:   "pipelines",
		Short: "Manage Conduit Pipelines",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	// 1. Create a new command to list pipelines
	// 2. Add that command as `pipelines ls` to the pipelineCmd

	pipelinesEditCmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit a pipeline",
		Run: func(cmd *cobra.Command, args []string) {
			openFile("./pipelines/generator-to-log.yaml")
		},
	}

	pipelinesListCmd := &cobra.Command{
		Use:     "ls",
		Aliases: []string{"list"},
		Short:   "List all pipelines",
		Run: func(cmd *cobra.Command, args []string) {
			//// TODO: Check if I'm running the service correctly
			//
			pipelinePath, _ := cmd.Flags().GetString("path")
			listDirectory(pipelinePath)
			//
			//logger := conduit.NewLogger(cfg.Log.Level, cfg.Log.Format)
			//db := &inmemory.DB{}
			//pipelineService := pipeline.NewService(logger, db)
			//
			//provisionService := provisioning.NewService(db, logger, pipelineService, nil, nil, nil, pipelinePath)
			//
			//r := &conduit.Runtime{
			//	Config:           cfg,
			//	DB:               db,
			//	Ready:            make(chan struct{}),
			//	Logger:           logger,
			//	ProvisionService: provisionService,
			//	PipelineService:  pipelineService,
			//}
			//r.ProvisionService.Init(ctx)
			//
			//// TODO: Init without provision
			//
			//r.PipelineService.Init(ctx)
			//r.PipelineService.Run(ctx, nil, nil, nil)
			//
			//instances := r.PipelineService.List(ctx)
			//for _, instance := range instances {
			//	fmt.Println(instance)
			//}
		},
	}
	pipelinesListCmd.PersistentFlags().StringVar(&cfg.Pipelines.Path, "path", cfg.Pipelines.Path, "path to the directory that has the yaml pipeline configuration files, or a single pipeline configuration file")

	pipelineCmd.AddCommand(pipelinesListCmd)
	pipelineCmd.AddCommand(cmdInit())
	pipelineCmd.AddCommand(pipelinesEditCmd)

	return pipelineCmd
}
