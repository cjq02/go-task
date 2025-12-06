package main

import "fmt"

// plusOne 将大整数数组加1
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		
		if digits[i] < 10 {
			return digits
		}
		
		digits[i] = 0
	}
	
	result := make([]int, len(digits)+1)
	result[0] = 1
	
	return result
}

func main() {
	digits1 := []int{1, 2, 3}
	result1 := plusOne(digits1)
	fmt.Printf("输入: %v\n", []int{1, 2, 3})
	fmt.Printf("输出: %v\n", result1)
	fmt.Printf("解释: 123 + 1 = 124\n")
	fmt.Println()
	
	digits2 := []int{4, 3, 2, 1}
	result2 := plusOne(digits2)
	fmt.Printf("输入: %v\n", []int{4, 3, 2, 1})
	fmt.Printf("输出: %v\n", result2)
	fmt.Printf("解释: 4321 + 1 = 4322\n")
	fmt.Println()
	
	digits3 := []int{9}
	result3 := plusOne(digits3)
	fmt.Printf("输入: %v\n", []int{9})
	fmt.Printf("输出: %v\n", result3)
	fmt.Printf("解释: 9 + 1 = 10\n")
	fmt.Println()
	
	digits4 := []int{9, 9, 9}
	result4 := plusOne(digits4)
	fmt.Printf("输入: %v\n", []int{9, 9, 9})
	fmt.Printf("输出: %v\n", result4)
	fmt.Printf("解释: 999 + 1 = 1000\n")
	fmt.Println()
	
	digits5 := []int{1, 9, 9}
	result5 := plusOne(digits5)
	fmt.Printf("输入: %v\n", []int{1, 9, 9})
	fmt.Printf("输出: %v\n", result5)
	fmt.Printf("解释: 199 + 1 = 200\n")
	fmt.Println()
	
	digits6 := []int{0}
	result6 := plusOne(digits6)
	fmt.Printf("输入: %v\n", []int{0})
	fmt.Printf("输出: %v\n", result6)
	fmt.Printf("解释: 0 + 1 = 1\n")
}

