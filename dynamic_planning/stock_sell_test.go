package dynamic_planning

import "testing"

func TestStockSell(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{arr: []int{3,1,0,0,4,5,8}},
			8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StockSell(tt.args.arr); got != tt.want {
				t.Errorf("StockSell() = %v, want %v", got, tt.want)
			}
		})
	}
}
