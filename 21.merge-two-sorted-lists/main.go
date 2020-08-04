package main

func main() {

}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode

	currentNode := &ListNode{}
	for l1 != nil || l2 != nil {
		node := &ListNode{}
		if l1.Val <= l2.Val {
			node.Val = l1.Val
			l1 = l1.Next
		} else {
			node.Val = l2.Val
			l2 = l2.Next
		}
		if head == nil {
			head = node
		} else {
			currentNode.Next = node
		}
		currentNode = node

	}
	return nil
}
