# Fiber 模板

一个基于 [Fiber](https://gofiber.io/) 框架的 Go Web 应用模板，采用清晰的分层架构和模块化设计。

## 特性

- 🚀 基于高性能的 Fiber 框架
- 📦 清晰的分层架构（控制器、服务、模型）
- 🔒 内置 JWT 认证
- 📝 完整的用户 CRUD API
- ⏰ 支持定时任务
- 🔧 灵活的配置管理
- 📊 集成 GORM ORM
- 🔍 请求验证和数据清理
- 📁 文件上传支持
- 📚 完善的文档

## 快速开始

### 前提条件

- Go 1.16 或更高版本
- PostgreSQL 数据库

### 安装

1. 克隆仓库
```bash
git clone https://github.com/NextEraAbyss/fiber-template.git
cd fiber-template
```

2. 安装依赖
```bash
go mod download
```

3. 配置环境变量
```bash
cp .env.example .env
# 编辑 .env 文件，配置数据库等信息
```

4. 运行应用
```bash
go run main.go
```

服务器将在 http://localhost:3000 启动（可通过 PORT 环境变量配置）。

## 项目结构

```
fiber-template/
├── app/                # 应用代码目录
│   ├── controller/     # 控制器目录
│   │   ├── auth_controller.go    # 认证相关控制器
│   │   └── user_controller.go    # 用户相关控制器
│   ├── middleware/     # 中间件目录
│   │   ├── auth.go              # 认证中间件
│   │   ├── cors.go              # CORS 中间件
│   │   ├── logger.go            # 日志中间件
│   │   └── validation.go        # 请求验证中间件
│   ├── model/          # 数据模型目录
│   │   ├── base.go              # 基础模型
│   │   └── user.go              # 用户模型
│   ├── public/         # 静态资源目录
│   │   └── uploads/            # 上传文件存储目录
│   ├── router/         # 路由目录
│   │   ├── v1/                 # API v1 版本路由
│   │   │   ├── auth.go         # 认证路由
│   │   │   └── user.go         # 用户路由
│   │   └── router.go           # 路由注册
│   ├── schedule/       # 定时任务目录
│   │   ├── loader.go           # 任务加载器
│   │   └── tasks/              # 具体任务实现
│   └── service/        # 服务层目录
│       ├── auth_service.go     # 认证服务
│       └── user_service.go     # 用户服务
├── config/             # 配置文件目录
│   ├── config.go              # 主配置文件
│   ├── database.go            # 数据库配置
│   ├── error.go               # 错误处理
│   ├── jwt.go                 # JWT 配置
│   ├── password.go            # 密码处理
│   ├── response.go            # 响应处理
│   ├── sanitizer.go           # 数据清理
│   ├── schedule.go            # 定时任务配置
│   └── validator.go           # 数据验证
├── docs/               # 文档目录
│   └── api/                   # API 文档
├── logs/               # 日志目录
├── test/               # 测试目录
│   ├── controller/            # 控制器测试
│   ├── service/              # 服务层测试
│   └── model/               # 模型测试
├── .env                # 环境变量配置
├── .env.example        # 环境变量示例
├── go.mod              # Go 模块定义
├── go.sum              # Go 依赖版本锁定
├── Makefile            # 项目管理命令
├── main.go             # 应用程序入口
└── README.md           # 项目文档
```

## 目录说明

### 应用代码 (`app/`)

- `controller/`: 处理 HTTP 请求和响应，调用相应的服务层方法
- `middleware/`: 提供请求预处理和后处理功能，如认证、日志记录等
- `model/`: 定义数据模型和数据库结构，包含字段验证规则
- `public/`: 存放静态文件和上传文件
- `router/`: 定义 API 路由，支持版本控制
- `schedule/`: 实现定时任务系统，支持 cron 表达式
- `service/`: 实现核心业务逻辑，处理数据操作

### 配置和工具 (`config/`)

- `config.go`: 主配置文件，包含环境变量加载
- `database.go`: 数据库配置和连接管理
- `jwt.go`: JWT 认证配置和工具
- `password.go`: 密码加密和验证
- `response.go`: HTTP 响应格式化
- `sanitizer.go`: 数据清理和验证
- `validator.go`: 请求数据验证
- `error.go`: 错误处理和定义
- `schedule.go`: 定时任务配置

### 其他目录

- `docs/`: 项目文档，包括 API 文档
- `logs/`: 应用日志文件
- `test/`: 单元测试和集成测试

## API 文档

API 文档位于 `docs/api/` 目录，包含以下内容：

- 认证接口
- 用户管理接口
- 文件上传接口
- 错误码说明

## 开发指南

### 添加新功能

1. 在 `app/model/` 中定义数据模型
2. 在 `app/service/` 中实现业务逻辑
3. 在 `app/controller/` 中处理 HTTP 请求
4. 在 `app/router/` 中注册路由

### 添加定时任务

1. 在 `app/schedule/tasks/` 中创建任务文件
2. 实现任务接口
3. 在 `app/schedule/loader.go` 中注册任务

### 添加中间件

1. 在 `app/middleware/` 中创建中间件文件
2. 实现中间件函数
3. 在 `app/router/router.go` 中注册中间件

## 环境变量

主要环境变量配置（完整列表见 `.env.example`）：

| 变量名 | 描述 | 默认值 |
|--------|------|--------|
| APP_NAME | 应用名称 | Fiber Template |
| APP_ENV | 运行环境 | development |
| PORT | 服务端口 | 3000 |
| DB_HOST | 数据库主机 | localhost |
| DB_PORT | 数据库端口 | 5432 |
| DB_USER | 数据库用户 | postgres |
| DB_PASSWORD | 数据库密码 | postgres |
| DB_NAME | 数据库名称 | fiber_template |
| JWT_SECRET | JWT 密钥 | your_jwt_secret_key |

## 贡献指南

1. Fork 项目
2. 创建特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建 Pull Request

## 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件 