package cli

import (
	"github.com/conduitio/conduit/pkg/conduit"
	"github.com/spf13/cobra"
)

func Run() {
	var cmdServe = &cobra.Command{
		Use:   "start",
		Short: "Start conduit service",
		Long:  `Starts the conduit service. This command will start the conduit service`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: add old arguments
			conduit.Serve(conduit.DefaultConfig())
		},
	}

	var rootCmd = &cobra.Command{Use: "conduit"}
	rootCmd.AddCommand(cmdServe)
	rootCmd.Execute()
}
