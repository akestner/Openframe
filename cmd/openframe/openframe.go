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
				cli.StringFlag{
					Name:   "config, c",
					Value:  openframe.ConfigDir(),
					Usage:  "Openframe.io config directory",
					EnvVar: "OPENFRAME_CONFIG",
				},
			},
			Action: func(c *cli.Context) {

				controllerConfig := openframe.ControllerOptions{
					Username:  c.String("username"),
					Password:  c.String("password"),
					Frame:     c.String("frame"),
					ConfigDir: c.String("config"),
				}

				frameController := openframe.Controller{}
				err := frameController.Init(controllerConfig)
				if err != nil {
					log.Fatal(err)
				}
			},
		},
	}

	app.Run(os.Args)
}
