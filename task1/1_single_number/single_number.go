package main

import "fmt"

// singleNumber 找出数组中只出现一次的数字
// 使用 map 记录每个元素出现的次数，然后遍历 map 找到出现次数为 1 的元素
func singleNumber(nums []int) int {
	// 使用 map 记录每个数字出现的次数
	countMap := make(map[int]int)
	
	// 使用 for 循环遍历数组，统计每个元素出现的次数
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		// 使用 if 条件判断，如果元素已存在则增加计数，否则初始化为 1
		if _, exists := countMap[num]; exists {
			countMap[num]++
		} else {
			countMap[num] = 1
		}
	}
	
	// 遍历 map，找到出现次数为 1 的元素
	for num, count := range countMap {
		// 使用 if 条件判断，找出出现次数为 1 的元素
		if count == 1 {
			return num
		}
	}
	
	// 如果没有找到，返回 0（根据题目假设，应该总是存在）
	return 0
}

func main() {
	// 测试用例 1: [2,2,1] 应该返回 1
	nums1 := []int{2, 2, 1}
	result1 := singleNumber(nums1)
	fmt.Printf("输入: %v\n", nums1)
	fmt.Printf("输出: %d\n", result1)
	fmt.Println()
	
	// 测试用例 2: [4,1,2,1,2] 应该返回 4
	nums2 := []int{4, 1, 2, 1, 2}
	result2 := singleNumber(nums2)
	fmt.Printf("输入: %v\n", nums2)
	fmt.Printf("输出: %d\n", result2)
	fmt.Println()
	
	// 测试用例 3: [1] 应该返回 1
	nums3 := []int{1}
	result3 := singleNumber(nums3)
	fmt.Printf("输入: %v\n", nums3)
	fmt.Printf("输出: %d\n", result3)
}

