#!/bin/bash

# 获取当前用户信息
# GET /api/users/me

source "$(dirname "$0")/../config.sh"

TOKEN=$(get_token)

if [ -z "$TOKEN" ]; then
    echo -e "${RED}错误: 未找到 token，请先登录${NC}"
    echo "使用: scripts/api/auth/login.sh <username> <password>"
    exit 1
fi

echo "获取当前用户信息"
echo ""

RESPONSE=$(curl -s -X GET "$BASE_URL/users/me" \
  -H "Authorization: Bearer $TOKEN")

print_response "$RESPONSE"

