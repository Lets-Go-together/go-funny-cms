package command

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func InitApp() *cli.App {
	app := &cli.App{
		Name:  "Start server",
		Usage: "--",
		Action: func(c *cli.Context) error {
			fmt.Println("Server")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "process",
				Aliases: []string{"s"},
				Usage:   "可以在这里触发一些自定义脚本",
				Action: func(c *cli.Context) error {
					fmt.Println("Ok")
					return nil
				},
			},
		},
	}

	return app
}
