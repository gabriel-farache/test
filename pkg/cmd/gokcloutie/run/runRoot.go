package run

import (
	"github.com/spf13/cobra"
	"github.com/IaC/go-kcloutie/pkg/cli"
	"github.com/IaC/go-kcloutie/pkg/params"
)

func Root(cliParams *params.Run, ioStreams *cli.IOStreams) *cobra.Command {
	cCmd := &cobra.Command{
		Use:     "run",
		Aliases: []string{},
		Short:   "Runs the web/api server",
	}
	cCmd.AddCommand(ServerCommand(cliParams, ioStreams))
	return cCmd
}
