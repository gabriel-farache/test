package version

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/IaC/go-kcloutie/pkg/cli"
	"github.com/IaC/go-kcloutie/pkg/cmd"
	"github.com/IaC/go-kcloutie/pkg/params/settings"
	"github.com/IaC/go-kcloutie/pkg/params/version"
	"gopkg.in/yaml.v3"
)

type VersionCmdOptions struct {
	IoStreams *cli.IOStreams
	CliOpts   *cli.CliOpts
	Output    string
}

func VersionCommand(ioStreams *cli.IOStreams) *cobra.Command {
	options := &VersionCmdOptions{}
	cCmd := &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   fmt.Sprintf("Prints %s version", settings.CliBinaryName),
		Run: func(cCmd *cobra.Command, args []string) {

			ctx := cmd.InitContextWithLogger("version", "")

			options.IoStreams = ioStreams

			options.CliOpts = cli.NewCliOptions()
			options.IoStreams.SetColorEnabled(!settings.RootOptions.NoColor)
			err := cmd.VerifyOutputParameterValue(options.Output)
			if err != nil {
				cmd.WriteCmdErrorToScreen(err.Error(), ioStreams, true, true)
			}
			cmd.CheckForUnknownArgsExitWhenFound(args, ioStreams)
			err = options.PrintVersion(ctx)
			if err != nil {
				cmd.WriteCmdErrorToScreen(err.Error(), ioStreams, true, true)
			}
		},
	}
	cCmd.PersistentFlags().StringVarP(&options.Output, "output", "o", "", "Output format. One of: (json, yaml)")
	return cCmd
}

func (o *VersionCmdOptions) PrintVersion(ctx context.Context) error {
	// log := logger.FromCtx(ctx)

	switch o.Output {
	case "":
		fmt.Fprintf(o.IoStreams.Out, "\n%s        %s\n", o.IoStreams.ColorScheme().GreenBold("Version:"), version.BuildVersion)
		fmt.Fprintf(o.IoStreams.Out, "%s         %s\n", o.IoStreams.ColorScheme().GreenBold("Commit:"), version.Commit)
		fmt.Fprintf(o.IoStreams.Out, "%s     %s\n\n", o.IoStreams.ColorScheme().GreenBold("Build Time:"), version.BuildTime)

	case "json":
		jsonByte, err := json.Marshal(newVersionDetails())
		if err != nil {
			cmd.WriteCmdErrorToScreen(fmt.Sprintf("failed to marshal JSON: %v", err), o.IoStreams, true, true)
		}
		fmt.Fprintf(o.IoStreams.Out, "%s\n", string(jsonByte))
	case "yaml":
		yamlByte, err := yaml.Marshal(newVersionDetails())
		if err != nil {
			cmd.WriteCmdErrorToScreen(fmt.Sprintf("failed to marshal YAML: %v", err), o.IoStreams, true, true)
		}
		fmt.Fprintf(o.IoStreams.Out, "%s\n", string(yamlByte))
	}
	return nil
}

func newVersionDetails() map[string]interface{} {

	return map[string]interface{}{
		"version":   version.BuildVersion,
		"commit":    version.Commit,
		"buildTime": version.BuildTime,
	}
}
