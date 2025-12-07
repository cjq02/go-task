#!/bin/bash

# 删除评论
# DELETE /api/comments/:id

source "$(dirname "$0")/../config.sh"

TOKEN=$(get_token)

if [ -z "$1" ]; then
    echo "用法: $0 <comment_id>"
    echo "示例: $0 1"
    exit 1
fi

COMMENT_ID="$1"

echo "删除评论 (ID: $COMMENT_ID)"
echo ""

if [ -z "$TOKEN" ]; then
    echo -e "${YELLOW}警告: 未找到 token，将调用 API 但不带认证信息${NC}"
    echo ""
    RESPONSE=$(curl -s -X DELETE "$BASE_URL/comments/$COMMENT_ID")
else
    RESPONSE=$(curl -s -X DELETE "$BASE_URL/comments/$COMMENT_ID" \
      -H "Authorization: Bearer $TOKEN")
fi

print_response "$RESPONSE"

