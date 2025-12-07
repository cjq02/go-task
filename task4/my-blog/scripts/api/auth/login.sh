#!/bin/bash

# 用户登录接口
# POST /api/auth/login

source "$(dirname "$0")/../config.sh"

if [ "$#" -lt 2 ]; then
    echo "用法: $0 <username> <password>"
    echo "示例: $0 testuser 123456"
    exit 1
fi

USERNAME="$1"
PASSWORD="$2"

echo "用户登录: $USERNAME"
echo ""

RESPONSE=$(curl -s -X POST "$BASE_URL/auth/login" \
  -H "Content-Type: application/json" \
  -d "{
    \"username\": \"$USERNAME\",
    \"password\": \"$PASSWORD\"
  }")

print_response "$RESPONSE"

# 保存 token 到文件
if command -v jq &> /dev/null; then
    TOKEN=$(echo "$RESPONSE" | jq -r '.data.token // empty')
else
    TOKEN=$(echo "$RESPONSE" | grep -o '"token":"[^"]*' | cut -d'"' -f4)
fi

if [ -n "$TOKEN" ] && [ "$TOKEN" != "null" ]; then
    save_token "$TOKEN"
    echo ""
    echo -e "${GREEN}Token 已保存到 $TOKEN_FILE${NC}"
fi

