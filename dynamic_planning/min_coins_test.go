package dynamic_planning

import "testing"

func TestMaxCoins1(t *testing.T) {
	type args struct {
		coins []int
		aim   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxCoins1(tt.args.coins, tt.args.aim); got != tt.want {
				t.Errorf("MaxCoins1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxCoins3(t *testing.T) {
	type args struct {
		coins []int
		aim   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				coins: []int{5,10,25,1},
				aim:   15,
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxCoins3(tt.args.coins, tt.args.aim); got != tt.want {
				t.Errorf("MaxCoins3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinCoins(t *testing.T) {
	type args struct {
		coins []int
		aim   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				coins: []int{2, 5, 7},
				aim:   27,
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinCoins(tt.args.coins, tt.args.aim); got != tt.want {
				t.Errorf("MinCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinCoins1(t *testing.T) {
	type args struct {
		coins []int
		aim   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				coins: []int{2, 5, 7},
				aim:   27,
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinCoins1(tt.args.coins, tt.args.aim); got != tt.want {
				t.Errorf("MinCoins1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_process1(t *testing.T) {
	type args struct {
		arr   []int
		index int
		aim   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := process1(tt.args.arr, tt.args.index, tt.args.aim); got != tt.want {
				t.Errorf("process1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxCoins4(t *testing.T) {
	type args struct {
		coins []int
		aim   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				coins: []int{5,10,25,1},
				aim:   15,
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxCoins4(tt.args.coins, tt.args.aim); got != tt.want {
				t.Errorf("MaxCoins4() = %v, want %v", got, tt.want)
			}
		})
	}
}