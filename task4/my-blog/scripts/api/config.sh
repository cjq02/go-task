#!/bin/bash

# API 测试脚本通用配置

BASE_URL="${BASE_URL:-http://localhost:9080/api}"
TOKEN_FILE="${TOKEN_FILE:-/tmp/blog_api_token.txt}"

# 颜色输出
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 保存 token 到文件
save_token() {
    echo "$1" > "$TOKEN_FILE"
}

# 从文件读取 token
get_token() {
    if [ -f "$TOKEN_FILE" ]; then
        cat "$TOKEN_FILE"
    fi
}

# 输出 JSON 响应
print_response() {
    if command -v jq &> /dev/null; then
        echo "$1" | jq .
    else
        echo "$1"
    fi
}

