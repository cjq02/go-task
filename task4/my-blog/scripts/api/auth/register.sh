#!/bin/bash

# 用户注册接口
# POST /api/auth/register

source "$(dirname "$0")/../config.sh"

if [ "$#" -lt 3 ]; then
    echo "用法: $0 <username> <password> <email>"
    echo "示例: $0 testuser 123456 test@example.com"
    exit 1
fi

USERNAME="$1"
PASSWORD="$2"
EMAIL="$3"

echo "注册用户: $USERNAME"
echo ""

RESPONSE=$(curl -s -X POST "$BASE_URL/auth/register" \
  -H "Content-Type: application/json" \
  -d "{
    \"username\": \"$USERNAME\",
    \"password\": \"$PASSWORD\",
    \"email\": \"$EMAIL\"
  }")

print_response "$RESPONSE"

