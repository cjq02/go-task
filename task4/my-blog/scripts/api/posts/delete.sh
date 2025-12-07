#!/bin/bash

# 删除文章
# DELETE /api/posts/:id

source "$(dirname "$0")/../config.sh"

TOKEN=$(get_token)

if [ -z "$1" ]; then
    echo "用法: $0 <post_id>"
    echo "示例: $0 1"
    exit 1
fi

POST_ID="$1"

echo "删除文章 (ID: $POST_ID)"
echo ""

if [ -z "$TOKEN" ]; then
    echo -e "${YELLOW}警告: 未找到 token，将调用 API 但不带认证信息${NC}"
    echo ""
    RESPONSE=$(curl -s -X DELETE "$BASE_URL/posts/$POST_ID")
else
    RESPONSE=$(curl -s -X DELETE "$BASE_URL/posts/$POST_ID" \
      -H "Authorization: Bearer $TOKEN")
fi

print_response "$RESPONSE"

