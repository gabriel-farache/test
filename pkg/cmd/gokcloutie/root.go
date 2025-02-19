package gokcloutie

import (
	"context"
	"fmt"

	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"github.com/IaC/go-kcloutie/pkg/cli"
	"github.com/IaC/go-kcloutie/pkg/cmd"
	"github.com/IaC/go-kcloutie/pkg/cmd/gokcloutie/run"
	"github.com/IaC/go-kcloutie/pkg/cmd/gokcloutie/version"
	"github.com/IaC/go-kcloutie/pkg/logger"
	"github.com/IaC/go-kcloutie/pkg/params"
	"github.com/IaC/go-kcloutie/pkg/params/settings"
)

var (
	showVersion = false
	ioStreams   = cli.NewIOStreams()
)

func Root(cliParams *params.Run) *cobra.Command {
	cCmd := &cobra.Command{
		Use:   "go-kcloutie",
		Short: "go-kcloutie is a cli/api tool for sending notifications",
		Long: heredoc.Doc(`
			go-kcloutie is a cli/api tool for sending notifications
		`),
		SilenceUsage: false,
		PersistentPreRun: func(cCmd *cobra.Command, args []string) {
			lgr := logger.Get()
			lgr.Info("Starting application")
			if settings.DebugModeEnabled || os.Getenv(settings.DebugModeLoggerEnvVar) != "" {
				lgr.Info("Debugging has been enabled!")
			}

		},
		RunE: func(cCmd *cobra.Command, args []string) error {
			if showVersion {
				vopts := version.VersionCmdOptions{
					IoStreams: ioStreams,
					CliOpts:   cli.NewCliOptions(),
					Output:    "",
				}
				vopts.IoStreams.SetColorEnabled(!settings.RootOptions.NoColor)
				vopts.PrintVersion(context.Background())
				return nil
			}
			return fmt.Errorf("no command was specified")
		},
		Annotations: map[string]string{
			"commandType": "main",
		},
	}

	cCmd.PersistentFlags().BoolVar(&settings.DebugModeEnabled, "debug", false, "When set, additional output around debugging is output to the screen")
	cCmd.PersistentFlags().BoolVarP(&settings.RootOptions.NoColor, cmd.NoColorFlag, "C", false, "Disable coloring")
	cCmd.PersistentFlags().BoolVar(&showVersion, "version", false, "Show the version")
	cCmd.AddCommand(version.VersionCommand(ioStreams))
	cCmd.AddCommand(run.Root(cliParams, ioStreams))

	return cCmd
}
