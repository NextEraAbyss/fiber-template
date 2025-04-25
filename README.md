# Fiber 模板

使用 Fiber 框架的 Go Web 应用模板，参考 Egg.js 的目录结构设计。

## 功能特点

- 清晰的模块化项目结构
- 带版本控制的路由系统
- 控制器、服务、模型分层架构
- 配置管理
- 错误处理
- 用户 CRUD API 示例
- JWT 身份认证
- PostgreSQL 数据库集成
- GORM ORM 支持
- 密码加密存储
- 中间件示例
- 定时任务系统

## 项目结构

```
fiber-template/
├── app/
│   ├── controller/     # 处理 HTTP 请求和响应
│   ├── middleware/     # 自定义中间件函数
│   ├── model/          # 数据模型和数据库架构
│   ├── router/         # 路由定义
│   ├── schedule/       # 定时任务
│   └── service/        # 业务逻辑
├── config/             # 配置文件
├── main.go             # 应用程序入口
└── README.md           # 项目文档
```

## 快速开始

### 前提条件

- Go 1.16 或更高版本
- PostgreSQL 数据库

### 安装

1. 克隆仓库
```bash
git clone https://github.com/yourusername/fiber-template.git
cd fiber-template
```

2. 安装依赖
```bash
go mod download
```

3. 配置数据库
确保你有一个可用的 PostgreSQL 数据库，并根据需要修改 `.env` 文件中的数据库配置。

4. 运行应用
```bash
go run main.go
```

服务器默认会在 3000 端口启动（可通过 PORT 环境变量配置）。

## API 端点

### 认证接口
- 登录: `POST /api/v1/auth/login`
- 注册: `POST /api/v1/auth/register`

### 用户接口 (需要JWT认证)
- 健康检查: `GET /api/health`
- 获取所有用户: `GET /api/v1/users`
- 获取单个用户: `GET /api/v1/users/:id`
- 创建用户: `POST /api/v1/users`
- 更新用户: `PUT /api/v1/users/:id`
- 删除用户: `DELETE /api/v1/users/:id`

## 定时任务

本项目支持Egg.js风格的定时任务系统，只需在`app/schedule`目录下创建文件即可。

### 创建定时任务

1. 在`app/schedule`目录下创建新的Go文件
2. 实现`Task`接口（Schedule、Task、Start、Stop方法）
3. 在`app/schedule/loader.go`的`LoadTasks()`函数中注册任务

```go
// 简单示例
package schedule

import (
	"log"
	"github.com/robfig/cron/v3"
)

type MyTask struct {
	cron *cron.Cron
}

func (t *MyTask) Schedule() string {
	return "*/30 * * * * *"  // cron表达式
}

func (t *MyTask) Task() {
	log.Println("执行我的任务")
}

// 其他必要方法...

func NewMyTask() *MyTask {
	return &MyTask{}
}
```

## 认证

本项目使用 JWT 进行认证。获取令牌后，在请求头中添加 `Authorization: Bearer <your_token>` 进行认证。

### 登录示例
```json
// POST /api/v1/auth/login
{
  "email": "user@example.com",
  "password": "password123"
}
```

### 注册示例
```json
// POST /api/v1/auth/register
{
  "username": "newuser",
  "email": "newuser@example.com",
  "password": "password123"
}
```

## 环境变量

| 变量名        | 描述                    | 默认值                  |
|--------------|------------------------|------------------------|
| APP_NAME     | 应用名称                | Fiber Template         |
| PORT         | HTTP 端口               | 3000                   |
| APP_ENV      | 环境（dev/prod）        | development            |
| DB_HOST      | 数据库主机              | localhost              |
| DB_PORT      | 数据库端口              | 5432                   |
| DB_USER      | 数据库用户              | postgres               |
| DB_PASSWORD  | 数据库密码              | password               |
| DB_NAME      | 数据库名称              | fiber_template         |
| JWT_SECRET   | JWT 签名密钥            | your_jwt_secret_key    |

## 许可证

本项目采用 MIT 许可证。 