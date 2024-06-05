package cli

import (
	"flag"
	"strings"
	"time"

	"github.com/conduitio/conduit/pkg/conduit"
	"github.com/spf13/cobra"
)

type StartFlags struct {
	StringValue  string        `long:"string-flag" usage:"A flag representing a string"`
	IntValue     int           `long:"int-flag" usage:"A flag representing an integer"`
	Duration     time.Duration `long:"duration-flag" usage:"A flag representing a duration"`
	StringsSlice []string      `long:"string-slice-flag" usage:"A flag representing a slice of strings"`
	BoolValue    bool          `long:"bool-flag" persistent:"false" usage:"A flag representing a boolean value"`
}

func cmdStart() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start conduit service",
		Long:  `Starts the conduit service. This command will start the conduit service`,
		Run: func(cmd *cobra.Command, args []string) {
			e := &conduit.Entrypoint{}
			cfg := conduit.DefaultConfig()
			e.Serve(cfg)
		},
	}
}

func Flags(cfg *conduit.Config) *flag.FlagSet {

	// TODO extract flags from config struct rather than defining flags manually
	flags := flag.NewFlagSet("conduit", flag.ExitOnError)

	flags.StringVar(&cfg.DB.Type, "db.type", cfg.DB.Type, "database type; accepts badger,postgres,inmemory")
	flags.StringVar(&cfg.DB.Badger.Path, "db.badger.path", cfg.DB.Badger.Path, "path to badger DB")
	flags.StringVar(&cfg.DB.Postgres.ConnectionString, "db.postgres.connection-string", cfg.DB.Postgres.ConnectionString, "postgres connection string")
	flags.StringVar(&cfg.DB.Postgres.Table, "db.postgres.table", cfg.DB.Postgres.Table, "postgres table in which to store data (will be created if it does not exist)")

	flags.BoolVar(&cfg.API.Enabled, "api.enabled", cfg.API.Enabled, "enable HTTP and gRPC API")
	flags.StringVar(&cfg.API.HTTP.Address, "http.address", cfg.API.HTTP.Address, "address for serving the HTTP API")
	flags.StringVar(&cfg.API.GRPC.Address, "grpc.address", cfg.API.GRPC.Address, "address for serving the gRPC API")

	flags.StringVar(&cfg.Log.Level, "log.level", cfg.Log.Level, "sets logging level; accepts debug, info, warn, error, trace")
	flags.StringVar(&cfg.Log.Format, "log.format", cfg.Log.Format, "sets the format of the logging; accepts json, cli")

	flags.StringVar(&cfg.Connectors.Path, "connectors.path", cfg.Connectors.Path, "path to standalone connectors' directory")
	flags.StringVar(&cfg.Processors.Path, "processors.path", cfg.Processors.Path, "path to standalone processors' directory")

	flags.StringVar(&cfg.Pipelines.Path, "pipelines.path", cfg.Pipelines.Path, "path to the directory that has the yaml pipeline configuration files, or a single pipeline configuration file")
	flags.BoolVar(&cfg.Pipelines.ExitOnError, "pipelines.exit-on-error", cfg.Pipelines.ExitOnError, "exit Conduit if a pipeline experiences an error while running")

	// NB: flags with prefix dev.* are hidden from help output by default, they only show up using '-dev -help'
	showDevHelp := flags.Bool("dev", false, "used together with the dev flag it shows dev flags")
	flags.StringVar(&cfg.Dev.CPUProfile, "dev.cpuprofile", "", "write cpu profile to file")
	flags.StringVar(&cfg.Dev.MemProfile, "dev.memprofile", "", "write memory profile to file")
	flags.StringVar(&cfg.Dev.BlockProfile, "dev.blockprofile", "", "write block profile to file")

	// show user or dev flags
	flags.Usage = func() {
		tmpFlags := flag.NewFlagSet("conduit", flag.ExitOnError)
		flags.VisitAll(func(f *flag.Flag) {
			if f.Name == "dev" || strings.HasPrefix(f.Name, "dev.") != *showDevHelp {
				return // hide flag from output
			}
			// reset value to its default, to ensure default is shown correctly
			_ = f.Value.Set(f.DefValue)
			tmpFlags.Var(f.Value, f.Name, f.Usage)
		})
		tmpFlags.Usage()
	}

	return flags
}
