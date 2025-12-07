#!/bin/bash

# 更新评论
# PUT /api/comments/:id

source "$(dirname "$0")/../config.sh"

TOKEN=$(get_token)

if [ -z "$TOKEN" ]; then
    echo -e "${RED}错误: 未找到 token，请先登录${NC}"
    echo "使用: scripts/api/auth/login.sh <username> <password>"
    exit 1
fi

if [ "$#" -lt 2 ]; then
    echo "用法: $0 <comment_id> <content>"
    echo "示例: $0 1 \"更新的评论内容\""
    exit 1
fi

COMMENT_ID="$1"
CONTENT="$2"

echo "更新评论 (ID: $COMMENT_ID)"
echo ""

RESPONSE=$(curl -s -X PUT "$BASE_URL/comments/$COMMENT_ID" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d "{
    \"content\": \"$CONTENT\"
  }")

print_response "$RESPONSE"

