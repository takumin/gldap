package client

import (
	"github.com/urfave/cli/v2"

	"github.com/takumin/gldap/internal/config"
	"github.com/takumin/gldap/internal/ldap/client"
)

func NewCommands(cfg *config.Config, flags []cli.Flag) *cli.Command {
	flags = append(flags, []cli.Flag{
		&cli.StringFlag{
			Name:        "endpoint-url",
			Aliases:     []string{"endpoint"},
			Usage:       "endpoint url",
			EnvVars:     []string{"ENDPOINT_URL", "ENDPOINT"},
			Value:       cfg.Client.Endpoint,
			Destination: &cfg.Client.Endpoint,
		},
	}...)
	return &cli.Command{
		Name:    "client",
		Aliases: []string{"c", "cli"},
		Usage:   "ldap client",
		Flags:   flags,
		Action:  action(cfg),
	}
}

func action(cfg *config.Config) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		cli, err := client.NewClient(ctx.Context, cfg)
		if err != nil {
			return err
		}
		return cli.Run()
	}
}
