package main

func main() {}

func shuffle(nums []int, n int) []int {
	ans := make([]int, 0, len(nums))
	for index, num := range nums {
		if index < n {
			tmpA := num
			tmpB := nums[index+n]
			ans = append(ans, tmpA, tmpB)
		}
	}
	return ans
}
