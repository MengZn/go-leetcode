package main

import (
	"reflect"
	"testing"
)

func Test_createTargetArray(t *testing.T) {
	type args struct {
		nums  []int
		index []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		//{
		//	name: "test1",
		//	args: args{
		//		nums:  []int{0, 1, 2, 3, 4},
		//		index: []int{0, 1, 2, 2, 1},
		//	},
		//	want: []int{0, 4, 1, 3, 2},
		//},
		{
			name: "test2",
			args: args{
				nums:  []int{4,2,1,1},
				index: []int{0,0,2,0},
			},
			want: []int{0,1,2,3,4},
		},
		{
			name: "test3",
			args: args{
				nums:  []int{1},
				index: []int{0},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createTargetArray(tt.args.nums, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createTargetArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
