package main

func rotate(nums []int, K int) {
	n := len(nums)
	if n == 0 || K%n == 0 {
		return
	}
	ans := make([]int, 2*n-K%n)
	for i := 0; i < len(ans); i++ {
		ans[i] = nums[i%n]
	}
	copy(nums, ans[len(ans)-n:])
}

func main() {
	nums := []int{9, 3, 7, 6, 4}
	rotate(nums, 3)
}
