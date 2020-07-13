package main

func main() {

}

func numIdenticalPairs(nums []int) int {
	positionMapper := make(map[int]int)
	count := 0

	for i := 0; i < len(nums); i++ {
		index := nums[i]
		for j := i; j < len(nums); j++ {
			pairindex := nums[j]
			if _, ok := positionMapper[i]; !ok {
				positionMapper[i] = index
				continue
			}
			if positionMapper[i] == pairindex && i < j {
				count++
			}
		}
	}
	return count
}
