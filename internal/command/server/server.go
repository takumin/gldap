package server

import (
	"github.com/urfave/cli/v2"

	"github.com/takumin/gldap/internal/config"
	"github.com/takumin/gldap/internal/ldap/server"
)

func NewCommands(c *config.Config, f []cli.Flag) *cli.Command {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "listen-url",
			Aliases:     []string{"listen"},
			Usage:       "listen url",
			EnvVars:     []string{"LISTEN_URL", "LISTEN"},
			Value:       c.Server.ListenURL,
			Destination: &c.Server.ListenURL,
		},
	}
	return &cli.Command{
		Name:    "server",
		Aliases: []string{"s", "serv"},
		Usage:   "ldap server",
		Flags:   append(flags, f...),
		Action:  action(c),
	}
}

func action(c *config.Config) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		cli, err := server.NewServer(ctx.Context, c)
		if err != nil {
			return err
		}
		return cli.Run()
	}
}
