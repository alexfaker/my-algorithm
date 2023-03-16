package arr_and_matrix

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetMaxSlidingWindow(t *testing.T) {
	type args struct {
		arr []int
		w   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"1",
			args{
				arr: []int{4,3,5,4,3,3,6,7},
				w:   3,
			},
			[]int{5,5,5,4,6,7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetMaxSlidingWindow(tt.args.arr, tt.args.w)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMaxSlidingWindow() = %v, want %v", got, tt.want)
			}
			fmt.Println(got)
		})
	}
}
