package main

func main() {}

func runningSum(nums []int) []int {
	current := 0
	ans := make([]int, len(nums))
	for idx, num := range nums {
		current = current + num
		ans[idx] = current
	}
	return ans
}
