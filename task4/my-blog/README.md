# 博客系统后端 API

基于 Gin 和 GORM 开发的博客系统后端服务。

## 功能特性

- 用户注册和登录（JWT 认证）
- 文章的 CRUD 操作
- 评论的 CRUD 操作
- 数据库自动迁移
- 日志记录

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
│   │   └── response.go      # 响应格式
│   └── logger/
│       └── logger.go        # 日志工具
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

