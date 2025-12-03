package main

import (
	"fmt"
	"sort"
)

// merge 合并所有重叠的区间
// 使用排序和遍历的方式，先按区间起始位置排序，然后合并重叠的区间
func merge(intervals [][]int) [][]int {
	// 如果区间数组为空或只有一个区间，直接返回
	if len(intervals) <= 1 {
		return intervals
	}

	// 按区间的起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 结果数组，初始化为第一个区间（创建新切片避免修改原数组）
	result := [][]int{{intervals[0][0], intervals[0][1]}}

	// 遍历排序后的区间数组
	for i := 1; i < len(intervals); i++ {
		// 获取当前区间
		current := intervals[i]
		// 获取结果数组中最后一个区间
		last := result[len(result)-1]

		// 判断当前区间是否与最后一个区间重叠
		// 如果当前区间的起始位置 <= 最后一个区间的结束位置，则重叠
		if current[0] <= last[1] {
			// 合并区间：取起始位置的最小值和结束位置的最大值
			// 起始位置：last[0]（因为已排序，last[0] <= current[0]）
			// 结束位置：取 last[1] 和 current[1] 的最大值
			if current[1] > last[1] {
				last[1] = current[1]
			}
			// 更新结果数组中的最后一个区间
			result[len(result)-1] = last
		} else {
			// 如果不重叠，将当前区间添加到结果数组（创建新切片避免修改原数组）
			result = append(result, []int{current[0], current[1]})
		}
	}

	return result
}

func main() {
	// 测试用例 1: [[1,3],[2,6],[8,10],[15,18]] -> [[1,6],[8,10],[15,18]]
	intervals1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	result1 := merge(intervals1)
	fmt.Printf("输入: %v\n", intervals1)
	fmt.Printf("输出: %v\n", result1)
	fmt.Printf("解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6]\n")
	fmt.Println()

	// 测试用例 2: [[1,4],[4,5]] -> [[1,5]]
	intervals2 := [][]int{{1, 4}, {4, 5}}
	result2 := merge(intervals2)
	fmt.Printf("输入: %v\n", intervals2)
	fmt.Printf("输出: %v\n", result2)
	fmt.Printf("解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间\n")
	fmt.Println()

	// 测试用例 3: [[4,7],[1,4]] -> [[1,7]]
	intervals3 := [][]int{{4, 7}, {1, 4}}
	result3 := merge(intervals3)
	fmt.Printf("输入: %v\n", intervals3)
	fmt.Printf("输出: %v\n", result3)
	fmt.Printf("解释: 区间 [1,4] 和 [4,7] 可被视为重叠区间\n")
	fmt.Println()

	// 测试用例 4: [[1,4],[0,4]] -> [[0,4]]
	intervals4 := [][]int{{1, 4}, {0, 4}}
	result4 := merge(intervals4)
	fmt.Printf("输入: %v\n", intervals4)
	fmt.Printf("输出: %v\n", result4)
	fmt.Printf("解释: 区间 [0,4] 完全包含 [1,4]\n")
	fmt.Println()

	// 测试用例 5: [[1,4],[2,3]] -> [[1,4]]
	intervals5 := [][]int{{1, 4}, {2, 3}}
	result5 := merge(intervals5)
	fmt.Printf("输入: %v\n", intervals5)
	fmt.Printf("输出: %v\n", result5)
	fmt.Printf("解释: 区间 [2,3] 完全包含在 [1,4] 中\n")
	fmt.Println()

	// 测试用例 6: [[1,4]] -> [[1,4]]
	intervals6 := [][]int{{1, 4}}
	result6 := merge(intervals6)
	fmt.Printf("输入: %v\n", intervals6)
	fmt.Printf("输出: %v\n", result6)
	fmt.Printf("解释: 只有一个区间，无需合并\n")
}
