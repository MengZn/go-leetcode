package main

func main() {

}
func decompressRLElist(nums []int) []int {
	res := make([]int, 0)
	for index, _ := range nums {
		if 2*index >= len(nums) {
			break
		}
		f := nums[2*index]
		v := nums[2*index+1]
		for i := 0; i < f; i++ {
			res = append(res, v)
		}
	}
	return res
}
