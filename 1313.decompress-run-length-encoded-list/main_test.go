package main

import (
	"reflect"
	"testing"
)

func Test_decompressRLElist(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test1",
			args: args{
				[]int{
					1, 2, 3, 4,
				},
			},
			want: []int{2, 4, 4, 4},
		},
		{
			name: "Test2",
			args: args{
				[]int{
					1, 1, 2, 3,
				},
			},
			want: []int{1, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decompressRLElist(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decompressRLElist() = %v, want %v", got, tt.want)
			}
		})
	}
}
