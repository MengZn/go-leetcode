package main

func main() {

}
func createTargetArray(nums []int, index []int) []int {
	numsIndices := make([]int, len(nums))
	result := make([]int, len(nums))
	// precalculates absolute indexes
	for i, relativeIndex := range index {
		numsIndices[i] = relativeIndex
		if relativeIndex < i {
			for j:=0; j < i; j++ {
				if numsIndices[j] >= relativeIndex {
					numsIndices[j]++
				}
			}
		}
	}
	for i := range(result) {
		result[numsIndices[i]] = nums[i]
	}
	return result
}
