package cli

import (
	"strings"

	"github.com/conduitio/conduit/pkg/conduit"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func cmdStart() *cobra.Command {
	cfg := conduit.DefaultConfig()

	start := &cobra.Command{
		Use:   "start",
		Short: "Start Conduit",
		Long:  `Starts the conduit service. This command will start the conduit service`,
		Run: func(cmd *cobra.Command, args []string) {
			e := &conduit.Entrypoint{}
			cfg := conduit.DefaultConfig()
			e.Serve(cfg)
		},
	}
	setStartFlags(&cfg, start)
	return start
}

// TODO: Consolidate with global flags that are in entrypoint.go
func setStartFlags(cfg *conduit.Config, cmd *cobra.Command) {

	cmd.PersistentFlags().StringVar(&cfg.DB.Type, "db.type", cfg.DB.Type, "database type; accepts badger,postgres,inmemory")
	cmd.PersistentFlags().StringVar(&cfg.DB.Badger.Path, "db.badger.path", cfg.DB.Badger.Path, "path to badger DB")
	cmd.PersistentFlags().StringVar(&cfg.DB.Postgres.ConnectionString, "db.postgres.connection-string", cfg.DB.Postgres.ConnectionString, "postgres connection string")
	cmd.PersistentFlags().StringVar(&cfg.DB.Postgres.Table, "db.postgres.table", cfg.DB.Postgres.Table, "postgres table in which to store data (will be created if it does not exist)")

	cmd.PersistentFlags().BoolVar(&cfg.API.Enabled, "api.enabled", cfg.API.Enabled, "enable HTTP and gRPC API")
	cmd.PersistentFlags().StringVar(&cfg.API.HTTP.Address, "http.address", cfg.API.HTTP.Address, "address for serving the HTTP API")
	cmd.PersistentFlags().StringVar(&cfg.API.GRPC.Address, "grpc.address", cfg.API.GRPC.Address, "address for serving the gRPC API")

	cmd.PersistentFlags().StringVar(&cfg.Log.Level, "log.level", cfg.Log.Level, "sets logging level; accepts debug, info, warn, error, trace")
	cmd.PersistentFlags().StringVar(&cfg.Log.Format, "log.format", cfg.Log.Format, "sets the format of the logging; accepts json, cli")

	cmd.PersistentFlags().StringVar(&cfg.Connectors.Path, "connectors.path", cfg.Connectors.Path, "path to standalone connectors' directory")
	cmd.PersistentFlags().StringVar(&cfg.Processors.Path, "processors.path", cfg.Processors.Path, "path to standalone processors' directory")

	cmd.PersistentFlags().StringVar(&cfg.Pipelines.Path, "pipelines.path", cfg.Pipelines.Path, "path to the directory that has the yaml pipeline configuration files, or a single pipeline configuration file")
	cmd.PersistentFlags().BoolVar(&cfg.Pipelines.ExitOnError, "pipelines.exit-on-error", cfg.Pipelines.ExitOnError, "exit Conduit if a pipeline experiences an error while running")

	// NB: flags with prefix dev.* are hidden from help output by default, they only show up using '-dev -help'
	showDevHelp := cmd.PersistentFlags().Bool("dev", false, "used together with the dev flag it shows dev flags")
	cmd.PersistentFlags().StringVar(&cfg.Dev.CPUProfile, "dev.cpuprofile", "", "write cpu profile to file")
	cmd.PersistentFlags().StringVar(&cfg.Dev.MemProfile, "dev.memprofile", "", "write memory profile to file")
	cmd.PersistentFlags().StringVar(&cfg.Dev.BlockProfile, "dev.blockprofile", "", "write block profile to file")

	// show user or dev flags
	// show user or dev flags
	cmd.PersistentFlags().Usage = func() {
		tmpFlags := pflag.NewFlagSet("conduit", pflag.ExitOnError)
		cmd.PersistentFlags().VisitAll(func(f *pflag.Flag) {
			if f.Name == "dev" || strings.HasPrefix(f.Name, "dev.") != *showDevHelp {
				return // hide flag from output
			}
			// reset value to its default, to ensure default is shown correctly
			_ = f.Value.Set(f.DefValue)
			tmpFlags.VarP(f.Value, f.Name, "", f.Usage)
		})
		tmpFlags.Usage()
	}
}
