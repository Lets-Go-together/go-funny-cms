## Go-Funny-CMS

[![Go Report Card](https://goreportcard.com/badge/github.com/Lets-Go-together/go-funny-cms)](https://goreportcard.com/report/github.com/Lets-Go-together/go-funny-cms)

## 项目

### 项目目录整体划分

- app 核心目录 ，控制器、model 等都在这里

- bootstrap:  辅助函数

- config:  配置加载

- pkg:  配置加载service

- public:  静态文件

- resources: 资源文件/前端项目

- .air.toml:  监听

- .env: 配置

- .env.example: 辅助配置参考

### 前端后台

    # 参考至
    git clone https://github.com/vueComponent/ant-design-vue-pro.git
    
目录地址: `resource/front`
    
编译:
    
    # 项目根目录
    
    npm install
    npm run server
    
### 前端客户端

目录地址: `resource/views`
    

### 相关包备注

    github.com/cheggaaa/pb/v3 进度条
    
### 第一次操作

    cp .env.example .env
    go build main.go && ./main generate-jwt
    # 或者
    go run main.go generate-jwt
    :按提示操作
    
    # 开发环境配置
    # 安装air
    # 根目录运行
    air
    
### 进度

**11-21**: 

    jwt、响应、日志、db等操作基本完成
    
**11-24**

    配置jwt密钥生成
    
**11-25**

    # 创建账户
    > go run main.go create-admin-user -h
    > go run main.go create-admin-user  --account [你的账户名称]
    
    # 登陆
    > air
    curl --location --request POST '127.0.0.1:8082/api/admin/login' \
    --header 'Content-Type: application/x-www-form-urlencoded' \
    --data-urlencode 'account=chenf' \
    --data-urlencode 'password=123456'

**11-27**

    # 新增中间件
    
    # 新增配置
    # 默认秒
    # pkg/auth/jwt.go:25
    JWT_EXPIRE_AT=10
    
    # 前端登陆注销对接完成

**研究一下项目要写什么东西**

参考项目

- 

### 资源路由定义

！遵从这个范式

![资源路由定义](https://images2015.cnblogs.com/blog/1128628/201703/1128628-20170321171908565-765352970.png)
