package cli

import (
	"context"
	"os"

	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	var rootCmd = &cobra.Command{Use: "conduit"}

	rootCmd.AddCommand(cmdStart())      // start command
	rootCmd.AddCommand(cmdVersion())    // version command
	rootCmd.AddCommand(cmdPipelines())  // pipeline commands
	rootCmd.AddCommand(cmdConnectors()) // connector commands
	rootCmd.AddCommand(cmdProcessors()) // processors commands
	rootCmd.AddCommand(cmdInit())       // processors commands
	rootCmd.AddCommand(cmdDoctor())     // processors commands
	return rootCmd
}

func Run() {
	ctx := context.Background()

	rootCmd := Cmd()
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
