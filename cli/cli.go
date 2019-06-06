package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func Run() {
	info()
	commands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func info() {
	app.Name = "Levpay Release CLI"
	app.Usage = "Levpay's release app manager"
	app.Author = "Levpay"
	app.Version = "1.0.0"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "--new",
			Aliases: []string{"-n"},
			Usage:   "Release a new app",
			Action:  new,
		},
	}
}

func new(c *cli.Context) {
	if len(c.Args()) == 0 {
		fmt.Println("Missing app argument")
		fmt.Println(c.Command.Usage)
		return
	}
}
