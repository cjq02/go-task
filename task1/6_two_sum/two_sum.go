package main

import "fmt"

// twoSum 找出数组中两个数的和等于目标值的索引
// 使用 map 数据结构，key 为数组元素的值，value 为数组元素的索引
// 遍历数组，对于每个元素，检查 target - nums[i] 是否在 map 中
func twoSum(nums []int, target int) []int {
	// 使用 map 存储已遍历的元素值和索引的映射关系
	numMap := make(map[int]int)

	// 遍历数组
	for i := 0; i < len(nums); i++ {
		// 计算需要的另一个数
		complement := target - nums[i]

		// 检查 complement 是否在 map 中
		if index, exists := numMap[complement]; exists {
			// 如果存在，返回两个索引
			return []int{index, i}
		}

		// 如果不存在，将当前元素和索引存入 map
		numMap[nums[i]] = i
	}

	// 如果没有找到，返回空数组（根据题目假设，应该总是存在答案）
	return []int{}
}

func main() {
	// 测试用例 1: nums = [2,7,11,15], target = 9 -> [0,1]
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	result1 := twoSum(nums1, target1)
	fmt.Printf("输入: nums = %v, target = %d\n", nums1, target1)
	fmt.Printf("输出: %v\n", result1)
	fmt.Printf("解释: 因为 nums[%d] + nums[%d] == %d，返回 [%d, %d]\n", result1[0], result1[1], target1, result1[0], result1[1])
	fmt.Println()

	// 测试用例 2: nums = [3,2,4], target = 6 -> [1,2]
	nums2 := []int{3, 2, 4}
	target2 := 6
	result2 := twoSum(nums2, target2)
	fmt.Printf("输入: nums = %v, target = %d\n", nums2, target2)
	fmt.Printf("输出: %v\n", result2)
	fmt.Printf("解释: 因为 nums[%d] + nums[%d] == %d，返回 [%d, %d]\n", result2[0], result2[1], target2, result2[0], result2[1])
	fmt.Println()

	// 测试用例 3: nums = [3,3], target = 6 -> [0,1]
	nums3 := []int{3, 3}
	target3 := 6
	result3 := twoSum(nums3, target3)
	fmt.Printf("输入: nums = %v, target = %d\n", nums3, target3)
	fmt.Printf("输出: %v\n", result3)
	fmt.Printf("解释: 因为 nums[%d] + nums[%d] == %d，返回 [%d, %d]\n", result3[0], result3[1], target3, result3[0], result3[1])
	fmt.Println()

	// 测试用例 4: nums = [-1,-2,-3,-4,-5], target = -8 -> [2,4]
	nums4 := []int{-1, -2, -3, -4, -5}
	target4 := -8
	result4 := twoSum(nums4, target4)
	fmt.Printf("输入: nums = %v, target = %d\n", nums4, target4)
	fmt.Printf("输出: %v\n", result4)
	fmt.Println()

	// 测试用例 5: nums = [1,5,3,7,9], target = 10 -> [1,3]
	nums5 := []int{1, 5, 3, 7, 9}
	target5 := 10
	result5 := twoSum(nums5, target5)
	fmt.Printf("输入: nums = %v, target = %d\n", nums5, target5)
	fmt.Printf("输出: %v\n", result5)
	fmt.Println()
}

