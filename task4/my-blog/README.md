# 博客系统后端 API

基于 Gin 和 GORM 开发的博客系统后端服务。

## 功能特性

- 用户注册和登录（JWT 认证）
- 退出登录功能
- 文章的 CRUD 操作（权限控制：只能修改/删除自己的文章）
- 评论的 CRUD 操作（权限控制：只能修改/删除自己的评论）
- 数据库自动迁移
- 统一的错误处理和日志记录
- 完善的权限验证机制

## 技术栈

- Go 1.21+
- Gin Web 框架
- GORM ORM
- MySQL 数据库
- JWT 认证

## 项目结构

```
my-blog/
├── cmd/
│   └── app/
│       └── main.go          # 应用入口
├── internal/
│   ├── app/
│   │   └── router.go        # 路由配置
│   ├── config/
│   │   ├── app.go           # 应用配置
│   │   └── database.go      # 数据库配置
│   ├── errors/
│   │   └── app_error.go     # 错误处理定义
│   ├── handler/
│   │   ├── user_handler.go  # 用户处理器
│   │   ├── post_handler.go  # 文章处理器
│   │   └── comment_handler.go # 评论处理器
│   ├── model/
│   │   ├── user.go          # 用户模型
│   │   ├── post.go          # 文章模型
│   │   └── comment.go       # 评论模型
│   ├── service/
│   │   ├── user_service.go  # 用户服务
│   │   ├── post_service.go  # 文章服务
│   │   ├── comment_service.go # 评论服务
│   │   └── jwt_service.go   # JWT 服务
│   ├── middleware/
│   │   ├── auth.go          # 认证中间件
│   │   └── cors.go          # CORS 中间件
│   ├── response/
│   │   └── response.go      # 响应格式和错误处理
│   └── logger/
│       └── logger.go        # 日志工具
├── scripts/
│   └── api/                 # API 测试脚本
│       ├── setup.sh         # 批量设置脚本执行权限
│       ├── config.sh        # 脚本配置文件
│       ├── auth/            # 认证相关脚本
│       ├── users/           # 用户相关脚本
│       ├── posts/           # 文章相关脚本
│       └── comments/        # 评论相关脚本
├── docs/                    # 文档目录
├── logs/                    # 日志目录
├── go.mod
└── README.md
```

## 环境配置

创建 `.env` 文件（参考 `.env.example`）：

```env
# 数据库配置
DB_HOST=localhost
DB_PORT=3306
DB_USER=task4_user
DB_PASSWORD=123456
DB_NAME=task4

# 应用配置
APP_ENV=development
APP_PORT=9080
JWT_SECRET_KEY=your-secret-key-change-in-production

# 日志配置
LOG_DIR=./logs
LOG_LEVEL=info
```

## 安装依赖

```bash
cd my-blog
go mod download
```

## 运行项目

```bash
cd my-blog/cmd/app
go run main.go
```

或者从项目根目录运行：

```bash
cd my-blog
go run cmd/app/main.go
```

## API 端点

### 认证相关

- `POST /api/auth/register` - 用户注册
- `POST /api/auth/login` - 用户登录
- `POST /api/auth/logout` - 退出登录（需要认证）

### 用户相关（需要认证）

- `GET /api/users/me` - 获取当前用户信息

### 文章相关

- `GET /api/posts` - 获取文章列表
- `GET /api/posts/:id` - 获取文章详情
- `POST /api/posts` - 创建文章（需要认证）
- `PUT /api/posts/:id` - 更新文章（需要认证）
- `DELETE /api/posts/:id` - 删除文章（需要认证）

### 评论相关

- `GET /api/comments/post/:postId` - 获取文章的所有评论
- `GET /api/comments/:id` - 获取评论详情
- `POST /api/comments` - 创建评论（需要认证）
- `PUT /api/comments/:id` - 更新评论（需要认证）
- `DELETE /api/comments/:id` - 删除评论（需要认证）

## 数据库迁移

项目启动时会自动运行数据库迁移，创建以下表：

- `users` - 用户表
- `posts` - 文章表
- `comments` - 评论表

## 错误处理

项目实现了统一的错误处理机制：

- **错误类型分类**：参数验证失败（400）、认证失败（401）、无权限（403）、资源不存在（404）、服务器错误（500）
- **统一的错误响应格式**：所有错误返回统一的 JSON 格式，包含错误码和错误消息
- **自动日志记录**：所有错误自动记录到日志文件，方便调试和维护

### HTTP 状态码映射

| 错误类型 | HTTP 状态码 | 业务错误码 |
|---------|------------|-----------|
| 参数验证失败 | 400 | 1001 |
| 认证失败 | 401 | 1002 |
| 无权限 | 403 | 1003 |
| 资源不存在 | 404 | 1004 |
| 数据库错误 | 500 | 2001 |
| 内部服务器错误 | 500 | 2002 |

### 错误响应示例

```json
{
  "code": 1003,
  "message": "无权限删除该文章或文章不存在"
}
```

## 权限控制

- **文章权限**：只有文章的作者才能更新或删除自己的文章
- **评论权限**：只有评论的作者才能更新或删除自己的评论
- **自动验证**：所有需要权限的操作都会自动验证用户身份

## 日志记录

- **日志级别**：INFO（正常操作）、WARN（客户端错误）、ERROR（服务器错误）
- **日志位置**：`logs/YYYY/MM/DD.log`，按日期自动分割
- **记录内容**：操作类型、用户ID、错误信息、请求路径等

## API 测试脚本

项目提供了完整的 API 测试脚本，位于 `scripts/api/` 目录：

### 快速开始

```bash
# 1. 设置脚本执行权限
chmod +x scripts/api/setup.sh
./scripts/api/setup.sh

# 2. 使用脚本测试 API
./scripts/api/auth/register.sh testuser 123456 test@example.com
./scripts/api/auth/login.sh testuser 123456
./scripts/api/posts/create.sh "文章标题" "文章内容"
```

详细使用方法请参考 [scripts/README.md](scripts/README.md)

## 测试

使用 curl 或其他 HTTP 客户端测试 API：

```bash
# 注册用户
curl -X POST http://localhost:9080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"123456","email":"test@example.com"}'

# 登录
curl -X POST http://localhost:9080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"123456"}'
```

或使用项目提供的测试脚本（推荐）：

```bash
./scripts/api/auth/register.sh testuser 123456 test@example.com
./scripts/api/auth/login.sh testuser 123456
```

