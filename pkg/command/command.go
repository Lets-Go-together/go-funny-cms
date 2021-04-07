package command

import (
	"github.com/urfave/cli/v2"
	"gocms/bootstrap"
	"gocms/example/pkg1"
	"gocms/pkg/auth"
	"gocms/pkg/auth/rabc"
	"gocms/pkg/casbin"
	"gocms/pkg/config"
	"gocms/pkg/mail/mailer"
	"gocms/pkg/pools"
	"gocms/pkg/schedule/backup"
)

func init() {
	config.Initialize()
	pools.Initialize()
	casbin.Initialize()
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
					pkg1.Echo()
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
				Name:    "schedule",
				Aliases: []string{"sch"},
				Usage:   "启动一个任务处理",
				Action: func(c *cli.Context) error {
					backup.InitSchedule()
					return nil
				},
			},
			{
				Name:    "schedule-exam",
				Aliases: []string{"sch-exam"},
				Usage:   "启动一个测试任务分发",
				Action: func(c *cli.Context) error {
					backup.DispatchTestProcess()
					return nil
				},
			},
			{
				Name:  "email-test",
				Usage: "测试邮件发送",
				Action: func(c *cli.Context) error {
					mailer := new(mailer.Mailer).LoadConfig()

					return nil
				},
			},
			{
				Name:    "generate-permission",
				Usage:   "初始化权限节点",
				Aliases: []string{"gp"},
				Action: func(c *cli.Context) error {
					bootstrap.Run(true)
					rabc.GeneratePermissionNodes()
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
		},
	}
}

// 服务器
func AppServer() {
	bootstrap.Run()
}
