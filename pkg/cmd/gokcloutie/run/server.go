package run

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"github.com/IaC/go-kcloutie/pkg/cli"
	"github.com/IaC/go-kcloutie/pkg/params/settings"
	"gopkg.in/yaml.v3"

	"github.com/IaC/go-kcloutie/pkg/api"
	"github.com/IaC/go-kcloutie/pkg/cmd"
	"github.com/IaC/go-kcloutie/pkg/params"
)

type ServerCmdOptions struct {
	IoStreams      *cli.IOStreams
	CliOpts        *cli.CliOpts
	ListeningAddr  string
	ConfigFilePath string
	CacheInSeconds int
}

func ServerCommand(run *params.Run, ioStreams *cli.IOStreams) *cobra.Command {
	options := &ServerCmdOptions{}
	cCmd := &cobra.Command{
		Use:     "server",
		Aliases: []string{"serv"},
		Short:   "Runs the API server",
		Long: heredoc.Docf(`
			Runs the API server on port 8080. To listen on a different port, use the %[1]s--listen-addr%[1]s flag. This command is blocking.

			If you do not include a %[1]s--config-file-path%[1]s for the API, then a basic default configuration is used.
		`, "`"),
		Example: heredoc.Doc(`
			# run an API server with a configuration
			go-kcloutie run server -c ./tests/files/serverConfig.json
		`),
		Run: func(cCmd *cobra.Command, args []string) {
			ctx := cmd.InitContextWithLogger("run", "server")
			serverConfig := api.NewServerConfiguration()
			if options.ConfigFilePath != "" {
				data, err := os.ReadFile(options.ConfigFilePath)
				if err != nil {
					cmd.WriteCmdErrorToScreen(fmt.Sprintf("failed to open configuration file - %v", err), ioStreams, true, true)
				}
				err = json.Unmarshal(data, serverConfig)
				if err != nil {
					err = yaml.Unmarshal(data, serverConfig)
					if err != nil {
						cmd.WriteCmdErrorToScreen(fmt.Sprintf("failed to unmarshal configuration file using json and yaml - %v", err), ioStreams, true, true)
					}
				}
			}

			ctx = api.WithCtx(ctx, serverConfig)

			options.IoStreams = ioStreams
			options.CliOpts = cli.NewCliOptions()
			options.IoStreams.SetColorEnabled(!settings.RootOptions.NoColor)
			cmd.CheckForUnknownArgsExitWhenFound(args, ioStreams)
			if cCmd.Flags().Changed("cache-expire-seconds") || serverConfig.CacheInSeconds == nil {
				serverConfig.CacheInSeconds = &options.CacheInSeconds
			}
			if cCmd.Flags().Changed("listen-addr") || serverConfig.ListeningAddr == nil {
				serverConfig.ListeningAddr = &options.ListeningAddr
			}
			err := serverConfig.Start(ctx)
			if err != nil {
				cmd.WriteCmdErrorToScreen(err.Error(), ioStreams, true, true)
			}
		},
	}
	cCmd.Flags().StringVarP(&options.ListeningAddr, "listen-addr", "l", ":8080", "The TCP address for the server to listen on, in the form \"host:port\". If empty, \":http\" (port 80) is used. The service names are defined in RFC 6335 and assigned by IANA. See net.Dial for details of the address format.")
	cCmd.Flags().StringVarP(&options.ConfigFilePath, "config-file-path", "c", "", "The path to the server configuration file")
	cCmd.Flags().IntVar(&options.CacheInSeconds, "cache-expire-seconds", 3600, "The number of seconds before cached values of the web server will expire")

	return cCmd
}
