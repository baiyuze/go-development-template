
## 🚀 Go 微服务框架模板（Gin + gRPC + Consul）

这是一个基于 Go 开发的微服务骨架，采用 Gin + gRPC + Consul + 自定义依赖注入容器 的架构，支持双协议（HTTP + gRPC）访问，适合构建高可维护的微服务系统。

⸻

## 🧱 项目结构总览
```
internal/
├── app_context/           # 上下文管理（如追溯ID等）
│   └── context.go         # 管理请求上下文的逻辑
├── common/                # 公共组件（如错误处理、日志）
│   ├── error/             # 错误处理
│   └── logger/            # 日志初始化与配置
│       └── logger.go      # 日志相关工具
├── container/             # 依赖注入容器（服务实例构造）
│   └── container.go       # 依赖注入初始化与管理
├── dto/                   # 数据传输对象（DTO），用于接口数据的转换
├── grpc/                  # gRPC 相关逻辑与配置
│   ├── client/            # gRPC 客户端封装（服务间调用）
│   │   └── user_client.go # gRPC 客户端实现
│   ├── container/         # gRPC 层注入（依赖管理）
│   │   └── container.go   # gRPC 服务所需依赖的注入
│   ├── handler/           # 实现 proto 定义的服务逻辑
│   │   └── hello_handler.go # 实现 gRPC 服务的逻辑
│   ├── proto/             # gRPC proto 文件及生成的代码
│   │   ├── hello.proto    # proto 文件定义
│   │   ├── hello.pb.go    # 通过 proto 生成的代码
│   │   └── hello_grpc.pb.go # 生成的 gRPC 服务代码
│   ├── register.go        # 注册 gRPC handler 到服务（可选）
│   └── server.go          # 启动 gRPC 服务（监听与注册）
├── handler/               # HTTP 请求处理逻辑
│   └── user.go            # 用户相关的 HTTP 处理器
├── middleware/            # 中间件逻辑（如日志、认证等）
├── model/                 # 数据模型层（用于数据库与数据结构定义）
├── repo/                  # 数据库交互层（用于数据持久化）
├── router/                # 路由注册
│   └── router.go          # 配置 Gin 路由与 gRPC 客户端
├── service/               # 核心业务逻辑（服务层）
│   └── user_service.go    # 用户相关的业务逻辑
├── tmp/                   # 临时目录（可以放缓存或日志）
├── utils/                 # 工具类（如 gRPC 工厂等）
│   ├── grpc_factory.go    # gRPC 客户端工厂
│   └── index.go           # 通用工具函数
├── .gitignore             # Git 忽略配置
├── build_proto.sh         # 编译 proto 文件的脚本
├── go.mod                 # Go 模块依赖
└── go.sum                 # Go 依赖的校验文件
```


⸻

## 🧪 启动方式

1. 安装依赖

首先，安装所有依赖：
```
go mod tidy
```
2. 启动服务

启动服务：
```
go run main.go
```
默认监听端口：
	•	HTTP 接口：http://localhost:8888
	•	gRPC 服务：localhost:50051

⸻

## 🧬 代码解析

main.go 入口文件
- 在 main.go 中，程序主要做了以下几项操作：
	- 环境变量与日志配置：
	  - 通过 os.Getenv("ENV") 获取当前环境（开发环境或生产环境），并相应配置 gin 的日志模式（开发模式或生产模式）。
		- 使用 zap 日志库，根据不同环境初始化不同的日志配置。
	- 中间件配置：
	  - 配置了一些常用的中间件，如：
	  - Trace：追溯请求 ID，用于日志追踪。
	  - AuthWhiteList：认证白名单，用于跳过特定的认证检查。
	  - Ginzap：在开发环境中输出日志到终端。
	  - RecoveryWithZap：自动恢复 panic 错误，并输出日志。
	  - Logger：自定义的日志中间件。
	- 配置初始化：
	  - 使用 config.InitConfig() 加载配置文件。如果加载失败，将会记录错误信息。
	- gRPC 服务启动：
	  - 通过 container.InitContainer() 初始化 gRPC 服务的依赖注入容器。
	  - 启动 gRPC 服务（server.IntServer()）。
	- gRPC 客户端初始化与路由注册：
	  - 使用 container.InitClient() 初始化 gRPC 客户端。
	  - 使用 router.RegisterRoutes() 注册 HTTP 路由。
	- 服务器运行：
	  - 最后，调用 r.Run(":8888") 启动 Gin HTTP 服务器。

⸻

## 🛰 gRPC 模块说明
```
模块名	说明
client/	封装 gRPC 客户端调用逻辑，支持复用连接、负载均衡等
container/	管理 handler 所需依赖，便于注入到 handler
handler/	实现 .proto 定义的服务接口，用于处理请求（如 SayHello）
proto/	包含 .proto 源文件及通过 protoc 生成的代码
register.go	把多个 handler 注册到 gRPC 服务中，保持 server.go 简洁
server.go	启动 gRPC 服务监听端口，注册健康检查、服务注册到 Consul 等
```


⸻

## 📦 RESTful API 示例（Gin）

GET 请求
```
http://localhost:8888/user/test
```
返回示例：
```
{
  "greeting": "你好, {user_data}!"
}
```


⸻

## 🧬 gRPC 接口测试

使用 grpcurl 命令行工具进行测试：
```
grpcurl -plaintext localhost:50051 app.HelloService.SayHello
```


⸻

## 🔧 环境变量支持

环境变量	描述	示例值
ENV	切换运行模式	development / production



⸻

## 📌 后续规划（TODO）
	•	添加用户鉴权中间件（JWT）

⸻

📄 License

MIT License.

⸻

其他说明

该框架结构清晰，易于扩展。通过 Gin 提供的路由与中间件支持 HTTP 服务，同时 gRPC 提供高效的微服务间通信。通过 Consul 提供服务发现与健康检查，提升系统的可维护性与可扩展性。依赖注入容器进一步解耦了各个模块，提高了代码的模块化和可测试性。
