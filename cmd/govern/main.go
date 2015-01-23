package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "govern"
	app.Usage = "A remote task execution environment"
	app.Version = "0.0.1"
	app.Author = "nesv"
	app.Email = "nicksaika <at> gmail <dot> com"

	app.Flags = []cli.Flag{}

	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "Run a playbook",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "i,inventory", Usage: "Path to an inventory file"},
			},
			Action: runPlaybook,
		},
	}

	app.Run(os.Args)
}
