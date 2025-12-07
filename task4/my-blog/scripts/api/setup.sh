#!/bin/bash

# 批量给所有 API 脚本添加执行权限

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"

echo "=== 批量设置 API 脚本执行权限 ==="
echo ""

if [ ! -d "$SCRIPT_DIR" ]; then
    echo "错误: 找不到当前目录: $SCRIPT_DIR"
    exit 1
fi

echo "正在查找所有 .sh 文件..."
echo ""

# 查找所有 .sh 文件并添加执行权限
count=0
while IFS= read -r -d '' file; do
    chmod +x "$file"
    rel_path="${file#$SCRIPT_DIR/}"
    echo "✓ $rel_path"
    count=$((count + 1))
done < <(find "$SCRIPT_DIR" -type f -name "*.sh" -print0)

echo ""
echo "✓ 已为 $count 个脚本设置执行权限！"
echo ""
echo "使用方法："
echo "  ./scripts/api/auth/register.sh <username> <password> <email>"
echo "  ./scripts/api/auth/login.sh <username> <password>"
echo "  ..."

