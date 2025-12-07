#!/bin/bash

# 退出登录接口
# POST /api/auth/logout

source "$(dirname "$0")/../config.sh"

TOKEN=$(get_token)

if [ -z "$TOKEN" ]; then
    echo -e "${RED}错误: 未找到 token，请先登录${NC}"
    echo "使用: scripts/api/auth/login.sh <username> <password>"
    exit 1
fi

echo "退出登录"
echo ""

RESPONSE=$(curl -s -X POST "$BASE_URL/auth/logout" \
  -H "Authorization: Bearer $TOKEN")

print_response "$RESPONSE"

# 清除 token
if [ -f "$TOKEN_FILE" ]; then
    rm "$TOKEN_FILE"
    echo ""
    echo -e "${GREEN}Token 已清除${NC}"
fi

