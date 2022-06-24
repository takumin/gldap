package server

import (
	"github.com/urfave/cli/v2"

	"github.com/takumin/gldap/internal/config"
	"github.com/takumin/gldap/internal/ldap/server"
)

func NewCommands(cfg *config.Config, flags []cli.Flag) *cli.Command {
	flags = append(flags, []cli.Flag{
		&cli.StringFlag{
			Name:        "listen-url",
			Aliases:     []string{"listen"},
			Usage:       "listen url",
			EnvVars:     []string{"LISTEN_URL", "LISTEN"},
			Value:       cfg.Server.ListenURL,
			Destination: &cfg.Server.ListenURL,
		},
	}...)
	return &cli.Command{
		Name:    "server",
		Aliases: []string{"s", "serv"},
		Usage:   "ldap server",
		Flags:   flags,
		Action:  action(cfg),
	}
}

func action(cfg *config.Config) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		cli, err := server.NewServer(ctx.Context, cfg)
		if err != nil {
			return err
		}
		return cli.Run()
	}
}
