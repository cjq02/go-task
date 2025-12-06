package main

import "fmt"

// singleNumber 找出数组中只出现一次的数字
func singleNumber(nums []int) int {
	countMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		countMap[num]++
	}

	for num, count := range countMap {
		if count == 1 {
			return num
		}
	}

	return 0
}

func main() {
	nums1 := []int{2, 2, 1}
	result1 := singleNumber(nums1)
	fmt.Printf("输入: %v\n", nums1)
	fmt.Printf("输出: %d\n", result1)
	fmt.Println()

	nums2 := []int{4, 1, 2, 1, 2}
	result2 := singleNumber(nums2)
	fmt.Printf("输入: %v\n", nums2)
	fmt.Printf("输出: %d\n", result2)
	fmt.Println()

	nums3 := []int{1}
	result3 := singleNumber(nums3)
	fmt.Printf("输入: %v\n", nums3)
	fmt.Printf("输出: %d\n", result3)
}
