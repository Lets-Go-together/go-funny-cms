## Golang 通用后台权限管理系统 (Go-Funny-CMS )

[![Go Report Card](https://goreportcard.com/badge/github.com/Lets-Go-together/go-funny-cms)](https://goreportcard.com/report/github.com/Lets-Go-together/go-funny-cms)


### 线上地址演示

    https://admin-go.surest.cn
    账号: surest
    密码: 123456

## 预览

![图片描述...](https://cdn.surest.cn/FuZGBSxkTbk_eS4OPM5FYqLZ6bQV)

## 项目地址

前端项目: https://github.com/Lets-Go-together/go-funny-cms-front
后端项目: https://github.com/Lets-Go-together/go-funny-cms


## 项目简介

是一个简单版本使用 `Casbin` + `Golang` 开发的通用后台权限管理系统

项目结构参考了`Laravel`初始化目录结构，更加便于 phper 进行开发和学习

目前采用的技术栈如下

- golang
- gin
- gorm(等)
- vue + design-vue
- casbin

采用前后端分离的开发方式

## 快速安装

    # 后端项目
    > https://github.com/Lets-Go-together/go-funny-cms.git
    > cd go-funny-cms
    > 导入sql: backups/funy_cms_20210514_153117.sql.gz
    > cp .env .env.example
    > go run main.go 
    # 或者
    > air

    # 前端项目
    > https://github.com/Lets-Go-together/go-funny-cms-front.git
    > cd go-funny-cms-front
    > yarn install
    > npm run dev

## 配置邮件发送

    # 后台运行
    > go run main.go express-run


## 额外命令

参考

    pkg/command/command.go

## 目前支持功能

- 后台账号管理
- 用户权限控制
- 自动权限路由生成
- RABC + ABC权限控制
- 自定义控制菜单栏
- 邮件发送与处理

## 目录结构

目前此系统未集成什么功能，非常便于二次开发进行，目录结构清晰

    - app :应用模块 （在次同级别目录，你可以同样创建app2目录）
        - http :api 接口操作相关
            - admin : 根据应用内模块区分
                - controler : 控制器层
                - validate : 关于reuqest 和 验证器都走这里
            - index : 例如客户端api 模块
                - 同上...
            - middleware : 用于中间件管理（可参考api 中间件的使用）

        - models : 模型
        - service: 字如其名 （service层）
        - validates: 验证器的二次封装
    - ... 中间的没什么好介绍的
    - pkg : 自定义创建的一些包，便于二次开发和提取

## 相关问题

登录下属用户的时候无菜单，请登录后台管理员，在菜单管理中 给对应用户的角色配置菜单

## 我的未来

由于时间的关系或者我个人的关系，需要去做一些更重要更值得做的事情，所以就草草的收尾了这个项目，欢迎提出有趣的想法和见解，我们一起来个思想碰撞，我也在致力于做一些自己的产品。

以上这个项目，如果有有趣的想法，欢迎一起讨论，再基础上继续开发

我们都"不止于此" ~

## 微信群

![图片描述...](https://cdn.surest.cn/FtdzkM_QTY9BMzGHFdpthI5Rh6Jg)

如码已过期，可以加我QQ 1562135624 备注 Golang
