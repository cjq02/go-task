#!/bin/bash

# 获取文章的所有评论
# GET /api/comments/post/:postId

source "$(dirname "$0")/../config.sh"

if [ -z "$1" ]; then
    echo "用法: $0 <post_id>"
    echo "示例: $0 1"
    exit 1
fi

POST_ID="$1"

echo "获取文章评论列表 (Post ID: $POST_ID)"
echo ""

RESPONSE=$(curl -s -X GET "$BASE_URL/comments/post/$POST_ID")

print_response "$RESPONSE"

