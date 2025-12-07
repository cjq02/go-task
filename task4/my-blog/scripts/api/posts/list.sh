#!/bin/bash

# 获取文章列表
# GET /api/posts

source "$(dirname "$0")/../config.sh"

LIMIT="${1:-10}"
OFFSET="${2:-0}"

echo "获取文章列表 (limit=$LIMIT, offset=$OFFSET)"
echo ""

RESPONSE=$(curl -s -X GET "$BASE_URL/posts?limit=$LIMIT&offset=$OFFSET")

print_response "$RESPONSE"

