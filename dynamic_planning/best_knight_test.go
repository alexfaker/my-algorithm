package dynamic_planning

import "testing"

func TestMinHp(t *testing.T) {
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
			args{
				arr: [][]int{
					{-2,-3,3},
					{-5,-10,1},
					{0,30,-5},
				},
			},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinHp(tt.args.arr); got != tt.want {
				t.Errorf("MinHp() = %v, want %v", got, tt.want)
			}
		})
	}
}
