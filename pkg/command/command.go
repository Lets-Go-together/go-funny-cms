package command

import (
	"github.com/urfave/cli/v2"
	"gocms/app/task/wyyMusic"
	"gocms/bootstrap"
	"gocms/pkg/auth"
	"gocms/pkg/config"
	"gocms/pkg/database"
	"gocms/pkg/pools"
)

func init() {
	config.Initialize()
	database.Initialize()
	pools.Initialize()
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
					//var wg sync.WaitGroup
					//for i := 0; i < 20; i++ {
					//	wg.Add(1)
					//	pools.PoolsExample(i, &wg)
					//}
					//wg.Wait()
					//adminService := service.AdminService{}
					//adminService.GetList()
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
