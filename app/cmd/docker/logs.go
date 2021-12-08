package docker

import (
	"github.com/paramah/ledo/app/modules/compose"
	"github.com/paramah/ledo/app/modules/context"
	"github.com/urfave/cli/v2"
)

var CmdComposeLogs = cli.Command{
	Name:        "logs",
	Aliases:     []string{"l"},
	Usage:       "logs from containers",
	Description: `Get fqn docker image defined as main service in config file`,
	Action:      RunComposeLogs,
}

func RunComposeLogs(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	compose.ExecComposerLogs(ctx, cmd.Args())
	return nil
}
