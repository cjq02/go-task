# API 测试脚本

每个接口对应一个独立的脚本文件，方便单独测试和调用。

## 目录结构

```
scripts/
├── api/
│   ├── config.sh                    # 通用配置文件
│   ├── auth/                        # 认证相关接口
│   │   ├── register.sh              # 用户注册
│   │   └── login.sh                 # 用户登录
│   ├── users/                       # 用户相关接口
│   │   └── me.sh                    # 获取当前用户信息
│   ├── posts/                       # 文章相关接口
│   │   ├── list.sh                  # 获取文章列表
│   │   ├── get.sh                   # 获取文章详情
│   │   ├── create.sh                # 创建文章
│   │   ├── update.sh                # 更新文章
│   │   └── delete.sh                # 删除文章
│   └── comments/                    # 评论相关接口
│       ├── list_by_post.sh          # 获取文章评论列表
│       ├── get.sh                   # 获取评论详情
│       ├── create.sh                # 创建评论
│       ├── update.sh                # 更新评论
│       └── delete.sh                # 删除评论
└── README.md
```

## 使用说明

### 1. 赋予执行权限

**方式一：使用自动设置脚本（推荐）**

```bash
# 先给 setup.sh 添加执行权限
chmod +x scripts/api/setup.sh

# 运行脚本，自动给所有 API 脚本添加执行权限
./scripts/api/setup.sh
```

**方式二：手动设置**

```bash
chmod +x scripts/api/**/*.sh
```

### 2. 配置

所有脚本共享 `scripts/api/config.sh` 配置文件：

- `BASE_URL`: API 基础地址，默认 `http://localhost:9080/api`
- `TOKEN_FILE`: Token 存储文件路径，默认 `/tmp/blog_api_token.txt`

可以通过环境变量覆盖：

```bash
export BASE_URL=http://localhost:9080/api
```

### 3. 使用示例

#### 认证相关

```bash
# 注册用户
./scripts/api/auth/register.sh testuser 123456 test@example.com

# 登录（会自动保存 token）
./scripts/api/auth/login.sh testuser 123456

# 退出登录（需要先登录，会自动清除 token）
./scripts/api/auth/logout.sh
```

#### 用户相关

```bash
# 获取当前用户信息（需要先登录）
./scripts/api/users/me.sh
```

#### 文章相关

```bash
# 获取文章列表
./scripts/api/posts/list.sh 10 0

# 获取文章详情
./scripts/api/posts/get.sh 1

# 创建文章（需要先登录）
./scripts/api/posts/create.sh "文章标题" "文章内容"

# 更新文章（需要先登录，只能更新自己的文章）
./scripts/api/posts/update.sh 1 "新标题" "新内容"

# 删除文章（需要先登录，只能删除自己的文章）
./scripts/api/posts/delete.sh 1
```

#### 评论相关

```bash
# 获取文章评论列表
./scripts/api/comments/list_by_post.sh 1

# 获取评论详情
./scripts/api/comments/get.sh 1

# 创建评论（需要先登录）
./scripts/api/comments/create.sh 1 "评论内容"

# 更新评论（需要先登录，只能更新自己的评论）
./scripts/api/comments/update.sh 1 "更新的评论内容"

# 删除评论（需要先登录，只能删除自己的评论）
./scripts/api/comments/delete.sh 1
```

## 完整测试流程示例

```bash
# 1. 注册用户
./scripts/api/auth/register.sh alice 123456 alice@example.com

# 2. 登录（保存 token）
./scripts/api/auth/login.sh alice 123456

# 3. 创建文章
./scripts/api/posts/create.sh "我的第一篇文章" "这是文章内容"

# 4. 获取文章列表
./scripts/api/posts/list.sh

# 5. 获取文章详情（假设文章 ID 为 1）
./scripts/api/posts/get.sh 1

# 6. 创建评论
./scripts/api/comments/create.sh 1 "这是一条评论"

# 7. 获取评论列表
./scripts/api/comments/list_by_post.sh 1
```

## Token 管理

登录后，token 会自动保存到 `$TOKEN_FILE`（默认 `/tmp/blog_api_token.txt`）。需要认证的接口会自动读取该文件中的 token。

如果需要清除 token：

```bash
rm /tmp/blog_api_token.txt
```

## 依赖要求

- `curl`: HTTP 请求工具
- `jq`: JSON 格式化工具（可选，但推荐安装）

安装 jq（Ubuntu/Debian）：

```bash
sudo apt-get install jq
```

## 注意事项

1. 使用前确保服务器已启动
2. 需要认证的接口必须先登录
3. 只能更新/删除自己创建的文章和评论
4. 脚本路径需要使用相对路径或绝对路径执行
