package main

import (
	"fmt"
)

// 构建大根堆，不停将堆顶元素弹出（将堆顶元素与最后一个元素对调，然后heapSize--，
// 使得堆顶元素不参与下一次大根堆化过程）每弹出一次并把最大元素放到堆顶后继续大根堆化，直到比目标值大的元素都被弹出（放到队堆尾）
func findKthLargest(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := 0; i < k-1; i++ {
		heapSize--
		nums[0], nums[heapSize] = nums[heapSize], nums[0]
		maxHeap(nums, 0, heapSize)
	}
	return nums[0]
}

func buildMaxHeap(nums []int, heapSize int) {
	//只需要从前面一半遍历即可，因为后续与左右子树比较，而左右子树必然在后面一半
	for i := heapSize / 2; i >= 0; i-- {
		maxHeap(nums, i, heapSize)
	}
}

func maxHeap(nums []int, i, heapSize int) {
	l, r, largest := 2*i+1, 2*i+2, i
	if l < heapSize && nums[l] > nums[largest] {
		largest = l
	}
	if r < heapSize && nums[r] > nums[largest] {
		largest = r
	}
	if largest != i {
		nums[largest], nums[i] = nums[i], nums[largest]
		maxHeap(nums, largest, heapSize)
	}
}

func main() {
	//nums := []int{3, 2, 3, 1, 2, 4, 5, 7, 6}
	nums := []int{3, 2, 1, 5, 6, 4}
	ans := findKthLargest(nums, 2)
	fmt.Println(ans)
}
