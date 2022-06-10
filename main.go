package main

import (
	"log"
	"os"

	"github.com/ayntgl/discordo/config"
	"github.com/ayntgl/discordo/ui"
	"github.com/urfave/cli/v2"
	"github.com/zalando/go-keyring"
)

func main() {
	cliApp := cli.NewApp()

	t, _ := keyring.Get(config.Name, "token")
	cliApp.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "token",
			Usage:   "The client authentication token.",
			Value:   t,
			Aliases: []string{"t"},
		},
		&cli.StringFlag{
			Name:    "config",
			Usage:   "The path to the configuration file.",
			Value:   config.DefaultPath(),
			Aliases: []string{"c"},
		},
	}

	cliApp.Action = func(ctx *cli.Context) error {
		c := config.New()
		err := c.Load(ctx.String("config"))
		if err != nil {
			return err
		}

		app := ui.NewApplication(ctx.String("token"), c)
		return app.Start()
	}

	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
