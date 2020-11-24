package command

import (
	"github.com/urfave/cli/v2"
	"gocms/app/task/wyyMusic"
	validateExample "gocms/app/validates/example"
	"gocms/bootstrap"
	"gocms/pkg/config"
	"gocms/pkg/database"
)

func init() {
	config.Initialize()
}

func InitApp() *cli.App {
	return &cli.App{
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
			{
				Name:    "test",
				Aliases: []string{"s"},
				Usage:   "可以在这里触发一些自定义脚本",
				Action: func(c *cli.Context) error {
					validateExample.Validate()
					return nil
				},
			},
			{
				Name:    "generate-jwt",
				Aliases: []string{"s"},
				Usage:   "初始化生成jwt密钥",
				Action: func(c *cli.Context) error {
					//auth.GerateSign()
					return nil
				},
			},
			{
				Name:    "wyy_music",
				Aliases: []string{"s"},
				Usage:   "网易云自动打卡听歌",
				Action: func(c *cli.Context) error {
					config.Initialize()
					database.Initialize()
					wyyMusic.Run()
					return nil
				},
			},
		},
	}
}

// 服务器
func AppServer() {
	database.Initialize()
	bootstrap.SteupRoute()
}
