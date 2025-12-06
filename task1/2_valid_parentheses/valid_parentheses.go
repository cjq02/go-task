package main

import "fmt"

// isValid 判断括号字符串是否有效
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if left, isRight := pairs[char]; isRight {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != left {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}

func main() {
	test1 := "()"
	result1 := isValid(test1)
	fmt.Printf("输入: \"%s\"\n", test1)
	fmt.Printf("输出: %v\n", result1)
	fmt.Println()

	test2 := "()[]{}"
	result2 := isValid(test2)
	fmt.Printf("输入: \"%s\"\n", test2)
	fmt.Printf("输出: %v\n", result2)
	fmt.Println()

	test3 := "(]"
	result3 := isValid(test3)
	fmt.Printf("输入: \"%s\"\n", test3)
	fmt.Printf("输出: %v\n", result3)
	fmt.Println()

	test4 := "([)]"
	result4 := isValid(test4)
	fmt.Printf("输入: \"%s\"\n", test4)
	fmt.Printf("输出: %v\n", result4)
	fmt.Println()

	test5 := "{[]}"
	result5 := isValid(test5)
	fmt.Printf("输入: \"%s\"\n", test5)
	fmt.Printf("输出: %v\n", result5)
	fmt.Println()

	test6 := "((("
	result6 := isValid(test6)
	fmt.Printf("输入: \"%s\"\n", test6)
	fmt.Printf("输出: %v\n", result6)
	fmt.Println()

	test7 := ")))"
	result7 := isValid(test7)
	fmt.Printf("输入: \"%s\"\n", test7)
	fmt.Printf("输出: %v\n", result7)
	fmt.Println()

	test8 := "(}"
	result8 := isValid(test8)
	fmt.Printf("输入: \"%s\"\n", test8)
	fmt.Printf("输出: %v\n", result8)
	fmt.Println()

	test9 := "((()))"
	result9 := isValid(test9)
	fmt.Printf("输入: \"%s\"\n", test9)
	fmt.Printf("输出: %v\n", result9)
}
