package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/conduitio/conduit/pkg/conduit"
	"github.com/spf13/cobra"
)

func cmdServe() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start conduit service",
		Long:  `Starts the conduit service. This command will start the conduit service`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: add old arguments
			conduit.Serve(conduit.DefaultConfig())
		},
	}
}

func cmdVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Conduit version",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: Show when was last updated
			_, _ = fmt.Fprintf(os.Stdout, "%s\n", conduit.Version(true))
		},
	}
}

func Cmd() *cobra.Command {
	var rootCmd = &cobra.Command{Use: "conduit"}

	// Add commands
	rootCmd.AddCommand(cmdServe(), cmdVersion())

	return rootCmd
}

func Run() {
	ctx := context.Background()

	rootCmd := Cmd()
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
