package completion

import (
	"github.com/urfave/cli/v2"

	"github.com/takumin/gldap/internal/command/completion/bash"
	"github.com/takumin/gldap/internal/command/completion/fish"
	"github.com/takumin/gldap/internal/command/completion/powershell"
	"github.com/takumin/gldap/internal/command/completion/zsh"
	"github.com/takumin/gldap/internal/config"
)

func NewCommands(c *config.Config, f []cli.Flag) *cli.Command {
	cmds := []*cli.Command{
		bash.NewCommands(c, f),
		fish.NewCommands(c, f),
		zsh.NewCommands(c, f),
		powershell.NewCommands(c, f),
	}
	return &cli.Command{
		Name:        "completion",
		Usage:       "command completion",
		Subcommands: cmds,
		HideHelp:    true,
	}
}
