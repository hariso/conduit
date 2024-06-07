package cli

import (
	"fmt"
	"os"

	"github.com/conduitio/conduit/pkg/conduit"
	"github.com/spf13/cobra"
)

func cmdDoctor() *cobra.Command {
	return &cobra.Command{
		Use:   "doctor",
		Short: "Check your conduit pipelines for potential problems a potential updates",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: Show when was last updated
			_, _ = fmt.Fprintf(os.Stdout, "%s\n", conduit.Version(true))
		},
	}
}
