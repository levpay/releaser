package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

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
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "Release a new app\t\t$ releaser n/new <app name>",
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
	err := os.Chdir(fmt.Sprintf("../%s", c.Args()[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Latest release of %s: ", c.Args()[0])
	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(string(stdout), err)
	}
	fmt.Println(string(stdout))
	question := fmt.Sprintf("Chose the new semantic version of the app %s: ", c.Args()[0])
	fmt.Print(question)
	reader := bufio.NewReader(os.Stdin)
	resp, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	if len(resp) == 0 {
		log.Fatal("response cannot be empty")
	}
	// resp = resp[:len(resp)-1]
	fmt.Print(resp)
	cmd = exec.Command("git", "tag", "-a", resp, "-m", "some message")
	stdout, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatal(string(stdout), err)
	}
	fmt.Println(string(stdout))
	// goreleaser --rm-dist --config=../releaser/goreleaser.yml
}
