# 错误处理与日志记录

本文档说明项目中的错误处理和日志记录机制。

## 错误处理架构

### 1. 统一的错误类型 (`internal/errors/app_error.go`)

定义了 `AppError` 结构，包含：
- `Code`: 业务错误码
- `HTTPStatus`: HTTP 状态码
- `Message`: 错误消息
- `Err`: 原始错误

### 2. 错误码定义

- `1001` - 参数验证失败 (400 Bad Request)
- `1002` - 认证失败 (401 Unauthorized)
- `1003` - 无权限 (403 Forbidden)
- `1004` - 资源不存在 (404 Not Found)
- `2001` - 数据库错误 (500 Internal Server Error)
- `2002` - 内部服务器错误 (500 Internal Server Error)

### 3. 错误处理函数

在 `internal/response/response.go` 中提供了：

- `HandleError()`: 统一处理错误，自动记录日志
- `HandleGormError()`: 专门处理 GORM 错误（如记录不存在）

## 日志记录

### 日志级别

- **INFO**: 正常操作（注册、登录、创建文章等）
- **WARN**: 客户端错误（认证失败、资源不存在等）
- **ERROR**: 服务器错误（数据库错误、内部错误等）

### 日志位置

日志文件保存在 `logs/` 目录下，按日期组织：
```
logs/
└── 2025/
    └── 12/
        └── 07.log
```

## 使用示例

### Handler 中的错误处理

```go
func (h *PostHandler) Create(c *gin.Context) {
    // 参数验证
    if err := c.ShouldBindJSON(&req); err != nil {
        appErr := errors.NewValidationError("请求参数无效: " + err.Error())
        response.HandleError(c, h.logger, appErr)
        return
    }

    // 业务逻辑
    post, err := h.postService.Create(userID, &req)
    if err != nil {
        response.HandleError(c, h.logger, err)
        return
    }

    // 记录日志
    h.logger.Info("用户 %d 创建了文章，ID: %d", userID, post.ID)
    response.Success(c, post.ToResponse())
}
```

### Service 中的错误返回

```go
func (s *UserService) Login(req *model.LoginRequest) (*model.User, error) {
    var user model.User
    if err := s.db.Where("username = ?", req.Username).First(&user).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, errors.NewAuthError("用户名或密码错误")
        }
        return nil, errors.NewDatabaseError(err)
    }
    // ...
}
```

## HTTP 状态码映射

| 错误类型 | HTTP 状态码 | 业务错误码 |
|---------|------------|-----------|
| 参数验证失败 | 400 | 1001 |
| 认证失败 | 401 | 1002 |
| 无权限 | 403 | 1003 |
| 资源不存在 | 404 | 1004 |
| 数据库错误 | 500 | 2001 |
| 内部服务器错误 | 500 | 2002 |

## 响应格式

### 成功响应
```json
{
  "code": 0,
  "data": {...}
}
```

### 错误响应
```json
{
  "code": 1001,
  "message": "请求参数无效"
}
```

## 已实现的改进

✅ 统一的错误处理机制
✅ 根据错误类型返回合适的 HTTP 状态码
✅ 自动日志记录（INFO/WARN/ERROR）
✅ 数据库连接错误处理
✅ 用户认证失败处理
✅ 资源不存在处理
✅ 权限检查错误处理

## 待完善的部分

- [ ] 更新所有 CommentHandler 方法使用新的错误处理
- [ ] 添加更多详细的日志信息
- [ ] 实现请求 ID 追踪

