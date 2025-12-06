package main

import "fmt"

// removeDuplicates 删除排序数组中的重复项
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	i := 0

	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			nums[i+1] = nums[j]
			i++
		}
	}

	return i + 1
}

func main() {
	nums1 := []int{1, 1, 2}
	k1 := removeDuplicates(nums1)
	fmt.Printf("输入: [1,1,2]\n")
	fmt.Printf("输出: %d, nums = %v\n", k1, nums1[:k1])
	fmt.Printf("解释: 函数返回新的长度 %d，前 %d 个元素为 %v\n", k1, k1, nums1[:k1])
	fmt.Println()

	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k2 := removeDuplicates(nums2)
	fmt.Printf("输入: [0,0,1,1,1,2,2,3,3,4]\n")
	fmt.Printf("输出: %d, nums = %v\n", k2, nums2[:k2])
	fmt.Printf("解释: 函数返回新的长度 %d，前 %d 个元素为 %v\n", k2, k2, nums2[:k2])
	fmt.Println()

	nums3 := []int{1}
	k3 := removeDuplicates(nums3)
	fmt.Printf("输入: [1]\n")
	fmt.Printf("输出: %d, nums = %v\n", k3, nums3[:k3])
	fmt.Println()

	nums4 := []int{1, 1, 1}
	k4 := removeDuplicates(nums4)
	fmt.Printf("输入: [1,1,1]\n")
	fmt.Printf("输出: %d, nums = %v\n", k4, nums4[:k4])
	fmt.Println()

	nums5 := []int{1, 2, 3, 4, 5}
	k5 := removeDuplicates(nums5)
	fmt.Printf("输入: [1,2,3,4,5]\n")
	fmt.Printf("输出: %d, nums = %v\n", k5, nums5[:k5])
	fmt.Println()

	nums6 := []int{-100, -100, 0, 0, 100, 100}
	k6 := removeDuplicates(nums6)
	fmt.Printf("输入: [-100,-100,0,0,100,100]\n")
	fmt.Printf("输出: %d, nums = %v\n", k6, nums6[:k6])
}
