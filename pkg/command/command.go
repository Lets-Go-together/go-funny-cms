package command

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"gocms/app/task/wyyMusic"
	"gocms/bootstrap"
	"gocms/pkg/auth"
	"gocms/pkg/config"
	"gocms/pkg/database"
	"strings"
)

func init() {
	config.Initialize()
	database.Initialize()
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
					//validateExample.Validate()
					r := strings.Split("Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MywiY3JlYXRlZF9hdCI6IjIwMjAtMTEtMjVUMDg6MjU6MzBaIiwidXBkYXRlZF9hdCI6IjIwMjAtMTEtMjVUMDg6MjU6MzBaIiwiYWNjb3VudCI6ImNoZW5mIiwicGFzc3dvcmQiOiIkMmEkMTQkczgxYlN6MnYxdDlJTzZEV01VOFhKZWdERzV1VjZBdVgvZy9mU1laL0k3U1JTMTB6bFNzSC4iLCJkZXNjcmlwdGlvbiI6IiIsImVtYWlsIjoiIiwicGhvbmUiOiIifQ.rwtCtzzhN2xWAcjRj2cqyK5u6RLyg1eiMy4YyOrb8Xo", "Bearer ")[0]
					fmt.Println(len(r))
					fmt.Println(r[1])
					return nil
				},
			},
			{
				Name:    "generate-jwt",
				Aliases: []string{"gj"},
				Usage:   "初始化生成jwt密钥",
				Action: func(c *cli.Context) error {
					auth.GerateSign()
					return nil
				},
			},
			{
				Name:    "create-admin-user",
				Aliases: []string{"cau"},
				Usage:   "创建一个admin用户",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "account",
						Value: "admin",
						Usage: "账户名称",
					},
				},
				Action: func(c *cli.Context) error {
					auth.GerateAdminUser(c.String("account"))
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
	bootstrap.SteupRoute()
}
