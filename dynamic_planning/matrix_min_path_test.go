package dynamic_planning

import "testing"

func getTestMatrix() [][]int{
	return [][]int{
		{1,3,5,9},
		{8,1,3,4},
		{5,0,6,1},
		{8,8,4,0},
	}
}
func TestGetMinPathSum(t *testing.T) {
	type args struct {
		arr [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{getTestMatrix()},
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMinPathSum(tt.args.arr); got != tt.want {
				t.Errorf("GetMinPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMinPathSum1(t *testing.T) {
	type args struct {
		arr [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{getTestMatrix()},
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMinPathSum1(tt.args.arr); got != tt.want {
				t.Errorf("GetMinPathSum1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMinPathSum2(t *testing.T) {
	type args struct {
		arr [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{getTestMatrix()},
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMinPathSum2(tt.args.arr); got != tt.want {
				t.Errorf("GetMinPathSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
