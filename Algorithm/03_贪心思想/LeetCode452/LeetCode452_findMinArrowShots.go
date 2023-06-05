package main

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	point := points[0][1]
	arrows := 1
	for i := 1; i < len(points); i++ {
		if points[i][0] > point {
			arrows++
			point = points[i][1]
		}
	}
	return arrows
}

func main() {
	points := [][]int{{2, 8}, {1, 6}, {7, 12}, {10, 16}}
	ans := findMinArrowShots(points)
	fmt.Println(ans)
}
