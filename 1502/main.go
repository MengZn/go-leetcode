package main

import "sort"

func main() {

}
func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	checker := arr[1] - arr[0]
	for i := 0; i < len(arr)-1; i++ {
		if arr[i]+checker != arr[i+1] {
			return false
		}
	}
	return true
}
