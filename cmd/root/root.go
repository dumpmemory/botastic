package root

import (
	"github.com/pandodao/botastic/cmd/app"
	"github.com/pandodao/botastic/cmd/gen"
	"github.com/pandodao/botastic/cmd/httpd"
	"github.com/pandodao/botastic/cmd/migrate"
	"github.com/pandodao/botastic/cmd/worker"
	"github.com/pandodao/botastic/config"
	"github.com/spf13/cobra"
)

func NewCmdRoot(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "botastic <command> <subcommand> [flags]",
		SilenceErrors: true,
		SilenceUsage:  true,
		Version:       version,
	}

	// load config
	config.C()

	cmd.AddCommand(httpd.NewCmdHttpd())
	cmd.AddCommand(migrate.NewCmdMigrate())
	cmd.AddCommand(gen.NewCmdGen())
	cmd.AddCommand(app.NewCmdApp())
	cmd.AddCommand(worker.NewCmdWorker())

	return cmd
}
