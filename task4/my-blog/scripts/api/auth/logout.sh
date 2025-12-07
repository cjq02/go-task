#!/bin/bash

# 退出登录接口
# POST /api/auth/logout

source "$(dirname "$0")/../config.sh"

TOKEN=$(get_token)

echo "退出登录"
echo ""

if [ -z "$TOKEN" ]; then
    echo -e "${YELLOW}警告: 未找到 token，将调用 API 但不带认证信息${NC}"
    echo ""
    RESPONSE=$(curl -s -X POST "$BASE_URL/auth/logout")
else
    RESPONSE=$(curl -s -X POST "$BASE_URL/auth/logout" \
      -H "Authorization: Bearer $TOKEN")
fi

print_response "$RESPONSE"

# 清除 token
if [ -f "$TOKEN_FILE" ]; then
    rm "$TOKEN_FILE"
    echo ""
    echo -e "${GREEN}Token 已清除${NC}"
fi

