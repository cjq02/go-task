#!/bin/bash

# 删除评论
# DELETE /api/comments/:id

source "$(dirname "$0")/../config.sh"

TOKEN=$(get_token)

if [ -z "$TOKEN" ]; then
    echo -e "${RED}错误: 未找到 token，请先登录${NC}"
    echo "使用: scripts/api/auth/login.sh <username> <password>"
    exit 1
fi

if [ -z "$1" ]; then
    echo "用法: $0 <comment_id>"
    echo "示例: $0 1"
    exit 1
fi

COMMENT_ID="$1"

echo "删除评论 (ID: $COMMENT_ID)"
echo ""

RESPONSE=$(curl -s -X DELETE "$BASE_URL/comments/$COMMENT_ID" \
  -H "Authorization: Bearer $TOKEN")

print_response "$RESPONSE"

