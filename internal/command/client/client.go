package client

import (
	"github.com/urfave/cli/v2"

	"github.com/takumin/gldap/internal/config"
	"github.com/takumin/gldap/internal/ldap/client"
)

func NewCommands(c *config.Config, f []cli.Flag) *cli.Command {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "endpoint-url",
			Aliases:     []string{"endpoint"},
			Usage:       "endpoint url",
			EnvVars:     []string{"ENDPOINT_URL", "ENDPOINT"},
			Value:       c.Client.Endpoint,
			Destination: &c.Client.Endpoint,
		},
	}
	return &cli.Command{
		Name:    "client",
		Aliases: []string{"c", "cli"},
		Usage:   "ldap client",
		Flags:   append(flags, f...),
		Action:  action(c),
	}
}

func action(c *config.Config) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		cli, err := client.NewClient(ctx.Context, c)
		if err != nil {
			return err
		}
		return cli.Run()
	}
}
