package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	current := head
	val := 0
	for current != nil {
		val = (val * 2) + current.Val
		current = current.Next
	}
	return val
}
