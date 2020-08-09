package main

import (
	"reflect"
	"testing"
)

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				&ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val:  -4,
								Next: nil,
							},
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceMakeLinkedList(t *testing.T) {
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

func sliceMakeLinkedList(slice []int) *ListNode {

	headNode := &ListNode{}
	tailNode := &ListNode{}
	for index, item := range slice {
		node := &ListNode{
			Val: item,
		}
		if index == 0 {
			headNode = node
		}
		tailNode.Next = node
		tailNode = node
	}
	return headNode
}
