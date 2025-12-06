package main

import "fmt"

// twoSum 找出数组中两个数的和等于目标值的索引
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]

		if index, exists := numMap[complement]; exists {
			return []int{index, i}
		}

		numMap[nums[i]] = i
	}

	return []int{}
}

func main() {
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	result1 := twoSum(nums1, target1)
	fmt.Printf("输入: nums = %v, target = %d\n", nums1, target1)
	fmt.Printf("输出: %v\n", result1)
	fmt.Printf("解释: 因为 nums[%d] + nums[%d] == %d，返回 [%d, %d]\n", result1[0], result1[1], target1, result1[0], result1[1])
	fmt.Println()

	nums2 := []int{3, 2, 4}
	target2 := 6
	result2 := twoSum(nums2, target2)
	fmt.Printf("输入: nums = %v, target = %d\n", nums2, target2)
	fmt.Printf("输出: %v\n", result2)
	fmt.Printf("解释: 因为 nums[%d] + nums[%d] == %d，返回 [%d, %d]\n", result2[0], result2[1], target2, result2[0], result2[1])
	fmt.Println()

	nums3 := []int{3, 3}
	target3 := 6
	result3 := twoSum(nums3, target3)
	fmt.Printf("输入: nums = %v, target = %d\n", nums3, target3)
	fmt.Printf("输出: %v\n", result3)
	fmt.Printf("解释: 因为 nums[%d] + nums[%d] == %d，返回 [%d, %d]\n", result3[0], result3[1], target3, result3[0], result3[1])
	fmt.Println()

	nums4 := []int{-1, -2, -3, -4, -5}
	target4 := -8
	result4 := twoSum(nums4, target4)
	fmt.Printf("输入: nums = %v, target = %d\n", nums4, target4)
	fmt.Printf("输出: %v\n", result4)
	fmt.Println()

	nums5 := []int{1, 5, 3, 7, 9}
	target5 := 10
	result5 := twoSum(nums5, target5)
	fmt.Printf("输入: nums = %v, target = %d\n", nums5, target5)
	fmt.Printf("输出: %v\n", result5)
	fmt.Println()
}

