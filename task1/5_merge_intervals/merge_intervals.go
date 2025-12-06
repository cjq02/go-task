package main

import (
	"fmt"
	"sort"
)

// merge 合并所有重叠的区间
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{{intervals[0][0], intervals[0][1]}}

	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		last := result[len(result)-1]

		if current[0] <= last[1] {
			if current[1] > last[1] {
				last[1] = current[1]
			}
			result[len(result)-1] = last
		} else {
			result = append(result, []int{current[0], current[1]})
		}
	}

	return result
}

func main() {
	intervals1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	result1 := merge(intervals1)
	fmt.Printf("输入: %v\n", intervals1)
	fmt.Printf("输出: %v\n", result1)
	fmt.Printf("解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6]\n")
	fmt.Println()

	intervals2 := [][]int{{1, 4}, {4, 5}}
	result2 := merge(intervals2)
	fmt.Printf("输入: %v\n", intervals2)
	fmt.Printf("输出: %v\n", result2)
	fmt.Printf("解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间\n")
	fmt.Println()

	intervals3 := [][]int{{4, 7}, {1, 4}}
	result3 := merge(intervals3)
	fmt.Printf("输入: %v\n", intervals3)
	fmt.Printf("输出: %v\n", result3)
	fmt.Printf("解释: 区间 [1,4] 和 [4,7] 可被视为重叠区间\n")
	fmt.Println()

	intervals4 := [][]int{{1, 4}, {0, 4}}
	result4 := merge(intervals4)
	fmt.Printf("输入: %v\n", intervals4)
	fmt.Printf("输出: %v\n", result4)
	fmt.Printf("解释: 区间 [0,4] 完全包含 [1,4]\n")
	fmt.Println()

	intervals5 := [][]int{{1, 4}, {2, 3}}
	result5 := merge(intervals5)
	fmt.Printf("输入: %v\n", intervals5)
	fmt.Printf("输出: %v\n", result5)
	fmt.Printf("解释: 区间 [2,3] 完全包含在 [1,4] 中\n")
	fmt.Println()

	intervals6 := [][]int{{1, 4}}
	result6 := merge(intervals6)
	fmt.Printf("输入: %v\n", intervals6)
	fmt.Printf("输出: %v\n", result6)
	fmt.Printf("解释: 只有一个区间，无需合并\n")
}
