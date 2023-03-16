package dynamic_planning

import (
	"fmt"
	"testing"
)

func TestLongestOrderArr(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"1",
			args{arr: []int{2, 1, 5, 3, 6, 4, 8, 9, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestOrderArr(tt.args.arr)
			fmt.Println(got)
		})
	}
}

func TestGetLongestSubArr(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"1",
			args{arr: []int{2, 1, 5, 3, 6, 4, 8, 9, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := GetLongestSubArr(tt.args.arr)
			fmt.Println(got)
		})
	}
}

func TestGetLongestSubArr2(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"1",
			args{arr: []int{2, 1, 5, 3, 6, 4, 8, 9, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetLongestSubArr2(tt.args.arr)
			fmt.Println(got)
		})
	}
}
