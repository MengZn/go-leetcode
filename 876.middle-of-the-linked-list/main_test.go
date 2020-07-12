package main

import (
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case1",
			args: args{
				sliceMakeSinglyLinkedList([]int{1, 2, 3, 4, 5}),
			},
			want: sliceMakeSinglyLinkedList([]int{3, 4, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sliceMakeSinglyLinkedList(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case1",
			args: args{
				[]int{1, 2, 3, 4},
			},
			want: &ListNode{
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
			if got := sliceMakeSinglyLinkedList(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sliceMakeSinglyLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func sliceMakeSinglyLinkedList(input []int) *ListNode {
	head := &ListNode{}
	tail := &ListNode{}

	head = nil
	tail = nil
	for _, value := range input {
		current := &ListNode{
			Val:  value,
			Next: nil,
		}
		if head == nil {
			head = current
		} else if head.Next == nil {
			tail = current
			head.Next = tail
		} else {
			tail.Next = current
			tail = current
		}

	}
	return head
}
