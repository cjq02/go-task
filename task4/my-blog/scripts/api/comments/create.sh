#!/bin/bash

# 创建评论
# POST /api/comments

source "$(dirname "$0")/../config.sh"

TOKEN=$(get_token)

if [ -z "$TOKEN" ]; then
    echo -e "${RED}错误: 未找到 token，请先登录${NC}"
    echo "使用: scripts/api/auth/login.sh <username> <password>"
    exit 1
fi

if [ "$#" -lt 2 ]; then
    echo "用法: $0 <post_id> <content>"
    echo "示例: $0 1 \"这是一条评论\""
    exit 1
fi

POST_ID="$1"
CONTENT="$2"

echo "创建评论 (Post ID: $POST_ID)"
echo ""

RESPONSE=$(curl -s -X POST "$BASE_URL/comments" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d "{
    \"postId\": $POST_ID,
    \"content\": \"$CONTENT\"
  }")

print_response "$RESPONSE"

