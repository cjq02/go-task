#!/bin/bash

# 获取评论详情
# GET /api/comments/:id

source "$(dirname "$0")/../config.sh"

if [ -z "$1" ]; then
    echo "用法: $0 <comment_id>"
    echo "示例: $0 1"
    exit 1
fi

COMMENT_ID="$1"

echo "获取评论详情 (ID: $COMMENT_ID)"
echo ""

RESPONSE=$(curl -s -X GET "$BASE_URL/comments/$COMMENT_ID")

print_response "$RESPONSE"

