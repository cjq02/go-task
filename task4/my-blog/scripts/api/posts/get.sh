#!/bin/bash

# 获取文章详情
# GET /api/posts/:id

source "$(dirname "$0")/../config.sh"

if [ -z "$1" ]; then
    echo "用法: $0 <post_id>"
    echo "示例: $0 1"
    exit 1
fi

POST_ID="$1"

echo "获取文章详情 (ID: $POST_ID)"
echo ""

RESPONSE=$(curl -s -X GET "$BASE_URL/posts/$POST_ID")

print_response "$RESPONSE"

