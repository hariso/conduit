package cli

import (
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
