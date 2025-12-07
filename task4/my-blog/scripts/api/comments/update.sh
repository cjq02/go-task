#!/bin/bash

# 更新评论
# PUT /api/comments/:id

source "$(dirname "$0")/../config.sh"

TOKEN=$(get_token)

if [ "$#" -lt 2 ]; then
    echo "用法: $0 <comment_id> <content>"
    echo "示例: $0 1 \"更新的评论内容\""
    exit 1
fi

COMMENT_ID="$1"
CONTENT="$2"

echo "更新评论 (ID: $COMMENT_ID)"
echo ""

if [ -z "$TOKEN" ]; then
    echo -e "${YELLOW}警告: 未找到 token，将调用 API 但不带认证信息${NC}"
    echo ""
    RESPONSE=$(curl -s -X PUT "$BASE_URL/comments/$COMMENT_ID" \
      -H "Content-Type: application/json" \
      -d "{
        \"content\": \"$CONTENT\"
      }")
else
    RESPONSE=$(curl -s -X PUT "$BASE_URL/comments/$COMMENT_ID" \
      -H "Content-Type: application/json" \
      -H "Authorization: Bearer $TOKEN" \
      -d "{
        \"content\": \"$CONTENT\"
      }")
fi

print_response "$RESPONSE"

