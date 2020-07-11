package main

import (
	"fmt"
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
	li := []int{1, 2, 3, 4}
	fmt.Printf("%v", makeHeadList(li))
}

func makeHeadList(input []int) *ListNode {
	head := &ListNode{}
	head = nil
	tail := &ListNode{}
	for _, val := range input {
		now := &ListNode{
			Val:  val,
			Next: nil,
		}
		if head == nil {
			head = now
			tail = now
		} else if head.Next == nil {
			head.Next = now
			tail = now
		} else if tail.Next == nil {
			tail.Next = now
			tail = now
		}
	}
	return head
}
