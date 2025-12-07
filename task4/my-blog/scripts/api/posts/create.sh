#!/bin/bash

# 创建文章
# POST /api/posts

source "$(dirname "$0")/../config.sh"

TOKEN=$(get_token)

if [ -z "$TOKEN" ]; then
    echo -e "${RED}错误: 未找到 token，请先登录${NC}"
    echo "使用: scripts/api/auth/login.sh <username> <password>"
    exit 1
fi

if [ "$#" -lt 2 ]; then
    echo "用法: $0 <title> <content>"
    echo "示例: $0 \"文章标题\" \"文章内容\""
    exit 1
fi

TITLE="$1"
CONTENT="$2"

echo "创建文章: $TITLE"
echo ""

RESPONSE=$(curl -s -X POST "$BASE_URL/posts" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d "{
    \"title\": \"$TITLE\",
    \"content\": \"$CONTENT\"
  }")

print_response "$RESPONSE"

