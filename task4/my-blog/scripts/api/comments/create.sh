#!/bin/bash

# 创建评论
# POST /api/comments

source "$(dirname "$0")/../config.sh"

TOKEN=$(get_token)

if [ "$#" -lt 2 ]; then
    echo "用法: $0 <post_id> <content>"
    echo "示例: $0 1 \"这是一条评论\""
    exit 1
fi

POST_ID="$1"
CONTENT="$2"

echo "创建评论 (Post ID: $POST_ID)"
echo ""

if [ -z "$TOKEN" ]; then
    echo -e "${YELLOW}警告: 未找到 token，将调用 API 但不带认证信息${NC}"
    echo ""
    RESPONSE=$(curl -s -X POST "$BASE_URL/comments" \
      -H "Content-Type: application/json" \
      -d "{
        \"postId\": $POST_ID,
        \"content\": \"$CONTENT\"
      }")
else
    RESPONSE=$(curl -s -X POST "$BASE_URL/comments" \
      -H "Content-Type: application/json" \
      -H "Authorization: Bearer $TOKEN" \
      -d "{
        \"postId\": $POST_ID,
        \"content\": \"$CONTENT\"
      }")
fi

print_response "$RESPONSE"

