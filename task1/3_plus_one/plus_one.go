package main

import "fmt"

// plusOne 将大整数数组加1
// 使用 for 循环从后往前遍历，处理进位情况
func plusOne(digits []int) []int {
	// 从数组末尾（最低位）开始遍历
	for i := len(digits) - 1; i >= 0; i-- {
		// 当前位加1
		digits[i]++
		
		// 如果当前位小于10，不需要进位，直接返回
		if digits[i] < 10 {
			return digits
		}
		
		// 如果当前位等于10，需要进位，当前位设为0
		digits[i] = 0
	}
	
	// 如果所有位都进位了（如 [9,9,9] -> [1,0,0,0]）
	// 需要在数组前面添加1
	result := make([]int, len(digits)+1)
	result[0] = 1
	// 其余位都是0，Go 的 make 已经初始化为0，所以不需要额外赋值
	
	return result
}

func main() {
	// 测试用例 1: [1,2,3] -> [1,2,4]
	digits1 := []int{1, 2, 3}
	result1 := plusOne(digits1)
	fmt.Printf("输入: %v\n", []int{1, 2, 3})
	fmt.Printf("输出: %v\n", result1)
	fmt.Printf("解释: 123 + 1 = 124\n")
	fmt.Println()
	
	// 测试用例 2: [4,3,2,1] -> [4,3,2,2]
	digits2 := []int{4, 3, 2, 1}
	result2 := plusOne(digits2)
	fmt.Printf("输入: %v\n", []int{4, 3, 2, 1})
	fmt.Printf("输出: %v\n", result2)
	fmt.Printf("解释: 4321 + 1 = 4322\n")
	fmt.Println()
	
	// 测试用例 3: [9] -> [1,0]
	digits3 := []int{9}
	result3 := plusOne(digits3)
	fmt.Printf("输入: %v\n", []int{9})
	fmt.Printf("输出: %v\n", result3)
	fmt.Printf("解释: 9 + 1 = 10\n")
	fmt.Println()
	
	// 测试用例 4: [9,9,9] -> [1,0,0,0]
	digits4 := []int{9, 9, 9}
	result4 := plusOne(digits4)
	fmt.Printf("输入: %v\n", []int{9, 9, 9})
	fmt.Printf("输出: %v\n", result4)
	fmt.Printf("解释: 999 + 1 = 1000\n")
	fmt.Println()
	
	// 测试用例 5: [1,9,9] -> [2,0,0]
	digits5 := []int{1, 9, 9}
	result5 := plusOne(digits5)
	fmt.Printf("输入: %v\n", []int{1, 9, 9})
	fmt.Printf("输出: %v\n", result5)
	fmt.Printf("解释: 199 + 1 = 200\n")
	fmt.Println()
	
	// 测试用例 6: [0] -> [1]
	digits6 := []int{0}
	result6 := plusOne(digits6)
	fmt.Printf("输入: %v\n", []int{0})
	fmt.Printf("输出: %v\n", result6)
	fmt.Printf("解释: 0 + 1 = 1\n")
}

