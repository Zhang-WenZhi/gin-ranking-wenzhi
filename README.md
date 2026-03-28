# go 版本

```shell
go version
# go version go1.20.2 windows/amd64
```

# 运行


```shell
# 检查redis是否启动
## 打开PowerShell 按 Win + X，选择 Windows PowerShell。
## 检查Redis服务状态 输入以下命令并按回车：
Get-Service | Where-Object {$_.Name -like "*redis*"}
# Stopped  redis              redis
```

```shell
# 只整理 go.mod，不升级版本
go mod tidy
# 更新某个指定依赖到最新版
# go get -u github.com/gin-gonic/gin
# 一键更新所有依赖到最新版本
# go get -u ./...
# 把依赖更新到指定版本
#go get github.com/gin-gonic/gin@v1.9.0
# 强制重新下载所有依赖
go mod download

# 1. main.go 右键运行
# 2. 对应目录 go run main.go
go run main.go
# 3.热重载（推荐）
# air 最新版要求 Go 1.22+，但你现在的 Go 版本太旧，不支持！
#go install github.com/air-verse/air@latest
go install github.com/cosmtrek/air@v1.49.0 
# go version go1.20.2 windows/amd64
# 终端运行命令，go env GOPATH的路径 + bin 才能看到ai r.exe，项目安装是get命令
# go env GOPATH的路径 + bin 加到环境变量
air -v
air
```

```shell
Another Redis Desktop Manager 客户端连接redis
# TODO 配置的redis windows 上没有，虚拟机上的，得linux上配合安装redis,修改redis 启动，才能访问：http://localhost:9999/ranking

# /player/list 可以访问！参数传入aid TODO: 日志没有SQL打印，后续加
```


# GoLang 编辑器，标红的库，检查实际有的就行，实际是能运行的，GoLang编辑器的bug

# gin-ranking-wenzhi

```shell
my-gin-app/
├── config/             # 配置文件目录
│   ├── config.go      # 配置结构体定义
│   └── database.go    # 数据库配置
├── controllers/        # 控制器目录，处理请求和响应
│   ├── user.go        # 用户相关控制器
│   └── product.go     # 产品相关控制器
├── middleware/         # 中间件目录
│   ├── auth.go        # 认证中间件
│   └── logger.go      # 日志中间件
├── models/            # 数据模型目录
│   ├── user.go        # 用户模型
│   └── product.go     # 产品模型
├── routes/            # 路由配置目录
│   └── routes.go      # 路由定义
├── services/          # 业务逻辑目录
│   ├── user.go        # 用户相关业务逻辑
│   └── product.go     # 产品相关业务逻辑
├── utils/             # 工具函数目录
│   ├── jwt.go         # JWT 工具
│   └── validator.go   # 验证工具
├── main.go            # 应用程序入口
└── go.mod             # 依赖管理文件

```

```shell
ffly-plus
├── config ## 配置
├── controller ## API实现,用来读取输入、调用业务处理、返回结果
│   └── api
│       └── v1
├── docs ## swag 文档
├── internal ##内部逻辑，业务目录
│   ├── cache　## 缓存
│   ├── code　## 错误码设计
│   ├── config　## 配置
│   ├── proto　## grpc proto
│   ├── sentinelm ## sentinel 限流
│   └── version　## 版本
├── models　## 数据库交互
├── pkg　## 一些封装好的 package
│   ├── token
│   └── utils
├── router ## 路由及中间件目录
│   ├── api
│   └── middleware
├── rpc ## 业务逻辑层
├── service ## 业务逻辑层
└── tool ## 小工具
    main.go # 项目入口文件
```