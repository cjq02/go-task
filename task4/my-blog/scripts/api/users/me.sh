#!/bin/bash

# 获取当前用户信息
# GET /api/users/me

source "$(dirname "$0")/../config.sh"

TOKEN=$(get_token)

echo "获取当前用户信息"
echo ""

if [ -z "$TOKEN" ]; then
    echo -e "${YELLOW}警告: 未找到 token，将调用 API 但不带认证信息${NC}"
    echo ""
    RESPONSE=$(curl -s -X GET "$BASE_URL/users/me")
else
    RESPONSE=$(curl -s -X GET "$BASE_URL/users/me" \
      -H "Authorization: Bearer $TOKEN")
fi

print_response "$RESPONSE"

