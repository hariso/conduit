package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/conduitio/conduit/pkg/conduit"
	"github.com/spf13/cobra"
)

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

	// start command
	start := cmdStart()
	cfg := conduit.DefaultConfig()
	SetStartFlags(&cfg, start)
	rootCmd.AddCommand(start)

	// version command
	rootCmd.AddCommand(cmdVersion())

	return rootCmd
}

func Run() {
	ctx := context.Background()

	rootCmd := Cmd()
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
