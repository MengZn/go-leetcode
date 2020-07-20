package main

import (
	"reflect"
	"testing"
)

func Test_getDecimalValue(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				makeHeadList([]int{1, 0, 1}),
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDecimalValue(tt.args.head); got != tt.want {
				t.Errorf("getDecimalValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_makeHeadList(t *testing.T) {
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
			if got := makeHeadList(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sliceMakeSinglyLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func makeHeadList(input []int) *ListNode {
	head := &ListNode{}
	tail := &ListNode{}

	head = nil
	tail = nil
	for _, val := range input {
		now := &ListNode{
			Val:  val,
			Next: nil,
		}
		if head == nil {
			head = now
			tail = now
		} else if head.Next == nil {
			tail = now
			head.Next = tail
		} else if tail.Next == nil {
			tail.Next = now
			tail = now
		}
	}
	return head
}
