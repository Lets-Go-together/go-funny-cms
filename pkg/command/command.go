package command

import (
	"github.com/urfave/cli/v2"
	"gocms/bootstrap"
)

func InitApp() *cli.App {
	app := &cli.App{
		Name:  "Start server",
		Usage: "--",
		Action: func(c *cli.Context) error {
			AppServer()
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "process",
				Aliases: []string{"s"},
				Usage:   "可以在这里触发一些自定义脚本",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
	}

	return app
}

// 服务器
func AppServer() {
	bootstrap.SteupRoute()
}
