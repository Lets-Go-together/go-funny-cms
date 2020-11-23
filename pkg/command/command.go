package command

import (
	"github.com/urfave/cli/v2"
	"gocms/app/task/wyyMusic"
	"gocms/bootstrap"
	"gocms/pkg/auth"
	"gocms/pkg/config"
	"gocms/pkg/database"
)

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
				Name:    "example-init",
				Aliases: []string{"s"},
				Usage:   "可以在这里触发一些自定义脚本",
				Action: func(c *cli.Context) error {
					//pkg1.Echo()
					//pkg2.Echo()
					return nil
				},
			},
			{
				Name:    "generate-jwt",
				Aliases: []string{"s"},
				Usage:   "初始化生成jwt密钥",
				Action: func(c *cli.Context) error {
					auth.Encrypto()
					return nil
				},
			},
			{
				Name: "wyy_music",
				Aliases: []string{"s"},
				Usage: "网易云自动打卡听歌",
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
	config.Initialize()
	database.Initialize()
	bootstrap.SteupRoute()
}
