package main

func main() {

}
func twoSum(nums []int, target int) []int {
	mapper := make(map[int]int, len(nums))
	for index, value := range nums {
		mapper[value] = index
	}

	for index, value := range nums {
		if reflectIndex, ok := mapper[target-value]; ok && reflectIndex != index{
			return []int{index, reflectIndex}
		}
	}
	return nil
}
