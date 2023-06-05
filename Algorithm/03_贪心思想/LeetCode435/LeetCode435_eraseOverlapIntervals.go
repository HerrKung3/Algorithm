package main

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 1 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	fmt.Println(intervals)
	//计数不重叠区间个数
	count := 0
	point := intervals[0][0]
	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] >= point {
			count++
			point = intervals[i][1]
		}
	}
	return len(intervals) - count
}

func main() {
	intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	ans := eraseOverlapIntervals(intervals)
	fmt.Println(ans)
}
