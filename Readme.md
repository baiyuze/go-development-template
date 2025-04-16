
# 🚀 Go 微服务框架模板（Gin + gRPC + Consul）

这是一个基于 Go 开发的微服务骨架，采用 **Gin + gRPC + Consul + 自定义依赖注入容器** 的架构，支持双协议（HTTP + gRPC）访问，适合构建高可维护的微服务系统。

---

## 🧱 项目结构总览
```
app/
├── config/              # 统一配置加载（支持 ENV 切换）
├── internal/
│   ├── container/       # 依赖注入容器（服务实例构造）
│   ├── grpc/
│   │   ├── client/      # gRPC 客户端封装（服务间调用）
│   │   ├── container/   # gRPC 层注入（用于注册时依赖注入）
│   │   ├── handler/     # 实现 proto 定义的服务逻辑
│   │   ├── proto/       # proto 文件及 gRPC 生成代码
│   │   ├── register.go  # 注册 gRPC handler 到服务（可选，职责分离）
│   │   └── server.go    # 启动 gRPC Server（服务注册、监听）
│   ├── router/          # Gin 路由注册（HTTP 接口）
│   └── service/         # 核心业务逻辑（service 层）
├── go.mod
└── main.go              # 应用启动入口
```
---

## 🧪 启动方式

### 1. 安装依赖

```bash
go mod tidy

2. 启动服务

go run main.go

默认监听端口：
	•	HTTP 接口：http://localhost:8888
	•	gRPC 服务：localhost:50051

⸻

🛰 gRPC 模块说明

模块名	说明
client/	封装 gRPC 客户端调用逻辑，支持复用连接、负载均衡等
container/	管理 handler 所需依赖，便于注入到 handler/
handler/	实现 .proto 定义的服务接口，用于处理请求（如 SayHello）
proto/	包含 .proto 源文件及通过 protoc 生成的代码
register.go	把多个 handler 注册到 gRPC 服务中，保持 server.go 简洁
server.go	启动 gRPC 服务监听端口，注册健康检查、服务注册到 Consul 等



⸻

📦 RESTful API 示例（Gin）

GET http://localhost:8888/hello

返回示例：

{
  "greeting": "你好, {user_data}!"
}



⸻

🧬 gRPC 接口测试

使用 grpcurl 命令行工具进行测试：

grpcurl -plaintext localhost:50051 app.HelloService.SayHello



⸻

🔧 环境变量支持

环境变量	描述	示例值
ENV	切换运行模式	development / production



⸻

📌 后续规划（TODO）
	•	添加用户鉴权中间件（JWT）
	•	服务拆分（User、Order、Auth 等）
	•	实现服务间 gRPC 调用示例
	•	添加 gRPC 客户端自动生成脚本
	•	引入服务注册与健康检查（Consul 加强版）

⸻

📄 License

MIT License.

---
