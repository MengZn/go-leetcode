package main

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case1",
			args: args{
				l1: sliceMakeLinkedList([]int{1, 2, 4}),
				l2: sliceMakeLinkedList([]int{1, 3, 4}),
			},
			want: sliceMakeLinkedList([]int{1, 1, 2, 3, 4, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_SliceMakeListNode(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name  string
		args  args
		wants *ListNode
	}{
		{
			name: "case1",
			args: args{
				[]int{1, 2, 3, 4},
			},
			wants: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sliceMakeLinkedList(tt.args.input); !reflect.DeepEqual(got, tt.wants) {
				t.Errorf("sliceMakeSinglyLinkedList() = %v, want %v", got, tt.wants)
			}
		})
	}
}

func sliceMakeLinkedList(input []int) *ListNode {
	var headNode *ListNode
	var tailNode *ListNode
	for _, value := range input {
		node := &ListNode{
			Val:  value,
			Next: nil,
		}
		if headNode == nil {
			headNode = node
			tailNode = node
			continue
		}
		if headNode.Next == nil {
			headNode.Next = node
			tailNode = node
		} else if tailNode.Next == nil {
			tailNode.Next = node
			tailNode = node
		}
	}
	return headNode
}
