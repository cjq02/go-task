#!/bin/bash

# 更新文章
# PUT /api/posts/:id

source "$(dirname "$0")/../config.sh"

TOKEN=$(get_token)

if [ "$#" -lt 3 ]; then
    echo "用法: $0 <post_id> <title> <content>"
    echo "示例: $0 1 \"新标题\" \"新内容\""
    exit 1
fi

POST_ID="$1"
TITLE="$2"
CONTENT="$3"

echo "更新文章 (ID: $POST_ID)"
echo ""

if [ -z "$TOKEN" ]; then
    echo -e "${YELLOW}警告: 未找到 token，将调用 API 但不带认证信息${NC}"
    echo ""
    RESPONSE=$(curl -s -X PUT "$BASE_URL/posts/$POST_ID" \
      -H "Content-Type: application/json" \
      -d "{
        \"title\": \"$TITLE\",
        \"content\": \"$CONTENT\"
      }")
else
    RESPONSE=$(curl -s -X PUT "$BASE_URL/posts/$POST_ID" \
      -H "Content-Type: application/json" \
      -H "Authorization: Bearer $TOKEN" \
      -d "{
        \"title\": \"$TITLE\",
        \"content\": \"$CONTENT\"
      }")
fi

print_response "$RESPONSE"

