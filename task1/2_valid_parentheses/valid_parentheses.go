package main

import "fmt"

// isValid 判断括号字符串是否有效
// 使用栈数据结构来处理括号匹配
//
// 规则要求：
// 1. 左括号必须用相同类型的右括号闭合
// 2. 左括号必须以正确的顺序闭合
// 3. 每个右括号都有一个对应的相同类型的左括号
func isValid(s string) bool {
	// 如果字符串长度为奇数，肯定不匹配（规则3：每个右括号都需要对应的左括号）
	if len(s)%2 != 0 {
		return false
	}

	// 使用切片模拟栈结构，用于存储未匹配的左括号
	stack := []rune{}

	// 定义括号映射关系：右括号 -> 对应的左括号
	// 规则1：确保相同类型的括号匹配（'(' 匹配 ')'，'{' 匹配 '}'，'[' 匹配 ']'）
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// 遍历字符串中的每个字符
	for _, char := range s {
		// 如果是右括号
		if left, isRight := pairs[char]; isRight {
			// 规则3：检查是否有对应的左括号
			// 如果栈为空，说明没有对应的左括号，返回 false
			if len(stack) == 0 {
				return false
			}
			// 规则1：检查栈顶的左括号是否与当前右括号类型相同
			// 规则2：通过栈的 LIFO 特性确保正确的闭合顺序（后进先出）
			if stack[len(stack)-1] != left {
				return false
			}
			// 匹配成功，弹出栈顶元素（表示该左括号已被正确闭合）
			stack = stack[:len(stack)-1]
		} else {
			// 如果是左括号，入栈（等待对应的右括号来闭合）
			stack = append(stack, char)
		}
	}

	// 规则3：如果栈为空，说明所有左括号都有对应的右括号
	// 如果栈不为空，说明还有未闭合的左括号，返回 false
	return len(stack) == 0
}

func main() {
	// 测试用例 1: "()" 应该返回 true
	test1 := "()"
	result1 := isValid(test1)
	fmt.Printf("输入: \"%s\"\n", test1)
	fmt.Printf("输出: %v\n", result1)
	fmt.Println()

	// 测试用例 2: "()[]{}" 应该返回 true
	test2 := "()[]{}"
	result2 := isValid(test2)
	fmt.Printf("输入: \"%s\"\n", test2)
	fmt.Printf("输出: %v\n", result2)
	fmt.Println()

	// 测试用例 3: "(]" 应该返回 false
	test3 := "(]"
	result3 := isValid(test3)
	fmt.Printf("输入: \"%s\"\n", test3)
	fmt.Printf("输出: %v\n", result3)
	fmt.Println()

	// 测试用例 4: "([)]" 应该返回 false（规则2：顺序错误）
	test4 := "([)]"
	result4 := isValid(test4)
	fmt.Printf("输入: \"%s\" (规则2测试：顺序错误)\n", test4)
	fmt.Printf("输出: %v\n", result4)
	fmt.Println()

	// 测试用例 5: "{[]}" 应该返回 true
	test5 := "{[]}"
	result5 := isValid(test5)
	fmt.Printf("输入: \"%s\"\n", test5)
	fmt.Printf("输出: %v\n", result5)
	fmt.Println()

	// 测试用例 6: "(((" 应该返回 false
	test6 := "((("
	result6 := isValid(test6)
	fmt.Printf("输入: \"%s\"\n", test6)
	fmt.Printf("输出: %v\n", result6)
	fmt.Println()

	// 测试用例 7: ")))" 应该返回 false（规则3：没有对应的左括号）
	test7 := ")))"
	result7 := isValid(test7)
	fmt.Printf("输入: \"%s\"\n", test7)
	fmt.Printf("输出: %v\n", result7)
	fmt.Println()

	// 测试用例 8: "(}" 应该返回 false（规则1：类型不匹配）
	test8 := "(}"
	result8 := isValid(test8)
	fmt.Printf("输入: \"%s\" (规则1测试：类型不匹配)\n", test8)
	fmt.Printf("输出: %v\n", result8)
	fmt.Println()

	// 测试用例 9: "((()))" 应该返回 true（规则2：正确顺序）
	test9 := "((()))"
	result9 := isValid(test9)
	fmt.Printf("输入: \"%s\" (规则2测试：正确顺序)\n", test9)
	fmt.Printf("输出: %v\n", result9)
}
