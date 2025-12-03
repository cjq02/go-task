package main

import "fmt"

// removeDuplicates 删除排序数组中的重复项
// 要求：原地修改输入数组，不使用额外的数组空间，O(1) 额外空间
// 使用双指针法：
//   - 慢指针 i：用于记录不重复元素的位置
//   - 快指针 j：用于遍历数组
//
// 当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位
func removeDuplicates(nums []int) int {
	// 如果数组为空或只有一个元素，直接返回长度
	if len(nums) <= 1 {
		return len(nums)
	}

	// 慢指针 i：用于记录不重复元素的位置
	i := 0

	// 快指针 j：用于遍历数组
	for j := 1; j < len(nums); j++ {
		// 当 nums[i] 与 nums[j] 不相等时
		if nums[i] != nums[j] {
			// 将 nums[j] 赋值给 nums[i + 1]
			nums[i+1] = nums[j]
			// 并将 i 后移一位
			i++
		}
		// 如果 nums[i] == nums[j]，说明是重复元素，j 继续向前遍历
	}

	// 返回唯一元素的数量（i 索引 + 1）
	return i + 1
}

func main() {
	// 测试用例 1: [1,1,2] -> 2, nums = [1,2,_]
	nums1 := []int{1, 1, 2}
	k1 := removeDuplicates(nums1)
	fmt.Printf("输入: [1,1,2]\n")
	fmt.Printf("输出: %d, nums = %v\n", k1, nums1[:k1])
	fmt.Printf("解释: 函数返回新的长度 %d，前 %d 个元素为 %v\n", k1, k1, nums1[:k1])
	fmt.Println()

	// 测试用例 2: [0,0,1,1,1,2,2,3,3,4] -> 5, nums = [0,1,2,3,4,_,_,_,_,_]
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k2 := removeDuplicates(nums2)
	fmt.Printf("输入: [0,0,1,1,1,2,2,3,3,4]\n")
	fmt.Printf("输出: %d, nums = %v\n", k2, nums2[:k2])
	fmt.Printf("解释: 函数返回新的长度 %d，前 %d 个元素为 %v\n", k2, k2, nums2[:k2])
	fmt.Println()

	// 测试用例 3: [1] -> 1, nums = [1]
	nums3 := []int{1}
	k3 := removeDuplicates(nums3)
	fmt.Printf("输入: [1]\n")
	fmt.Printf("输出: %d, nums = %v\n", k3, nums3[:k3])
	fmt.Println()

	// 测试用例 4: [1,1,1] -> 1, nums = [1,_,_]
	nums4 := []int{1, 1, 1}
	k4 := removeDuplicates(nums4)
	fmt.Printf("输入: [1,1,1]\n")
	fmt.Printf("输出: %d, nums = %v\n", k4, nums4[:k4])
	fmt.Println()

	// 测试用例 5: [1,2,3,4,5] -> 5, nums = [1,2,3,4,5]
	nums5 := []int{1, 2, 3, 4, 5}
	k5 := removeDuplicates(nums5)
	fmt.Printf("输入: [1,2,3,4,5]\n")
	fmt.Printf("输出: %d, nums = %v\n", k5, nums5[:k5])
	fmt.Println()

	// 测试用例 6: [-100,-100,0,0,100,100] -> 3, nums = [-100,0,100,_,_,_]
	nums6 := []int{-100, -100, 0, 0, 100, 100}
	k6 := removeDuplicates(nums6)
	fmt.Printf("输入: [-100,-100,0,0,100,100]\n")
	fmt.Printf("输出: %d, nums = %v\n", k6, nums6[:k6])
}
