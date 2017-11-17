package main

import (
	"github.com/akestner/openframe/openframe"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "openframe"
	app.Usage = "CLI interface to the `openframe` frame controller"
	app.Version = "0.4.9"

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "Run the Openframe Frame Controller",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "username, u",
					Usage:  "Openframe.io username",
					EnvVar: "OPENFRAME_USER",
				},
				cli.StringFlag{
					Name:   "password, p",
					Usage:  "Openframe.io password",
					EnvVar: "OPENFRAME_PASSWORD",
				},
				cli.StringFlag{
					Name:   "frame, f",
					Usage:  "Openframe.io frame name",
					EnvVar: "OPENFRAME_FRAME",
				},
			},
			Action: func(c *cli.Context) {
				controllerConfig := openframe.ControllerConfig{
					Username: c.String("username"),
					Password: c.String("password"),
					Frame:    c.String("frame"),
				}

				frameController := openframe.Controller{}
				if err := frameController.Init(controllerConfig); err != nil {
					log.Fatal(err)
				}
			},
		},
	}

	app.Run(os.Args)
}
