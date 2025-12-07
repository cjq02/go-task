#!/bin/bash

# 创建文章
# POST /api/posts

source "$(dirname "$0")/../config.sh"

TOKEN=$(get_token)

if [ "$#" -lt 2 ]; then
    echo "用法: $0 <title> <content>"
    echo "示例: $0 \"文章标题\" \"文章内容\""
    exit 1
fi

TITLE="$1"
CONTENT="$2"

echo "创建文章: $TITLE"
echo ""

if [ -z "$TOKEN" ]; then
    echo -e "${YELLOW}警告: 未找到 token，将调用 API 但不带认证信息${NC}"
    echo ""
    RESPONSE=$(curl -s -X POST "$BASE_URL/posts" \
      -H "Content-Type: application/json" \
      -d "{
        \"title\": \"$TITLE\",
        \"content\": \"$CONTENT\"
      }")
else
    RESPONSE=$(curl -s -X POST "$BASE_URL/posts" \
      -H "Content-Type: application/json" \
      -H "Authorization: Bearer $TOKEN" \
      -d "{
        \"title\": \"$TITLE\",
        \"content\": \"$CONTENT\"
      }")
fi

print_response "$RESPONSE"

