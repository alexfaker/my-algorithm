package arr_and_matrix

import "testing"


func getTestMatrix()[][]int{
	var matrix = [][]int{
		{1,2,3,4},
		{5,6,7,8},
		{9,10,11,12},
		{13,14,15,16},

	}
	return matrix
}
func TestPrintMatrixSpiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"1",
			args{getTestMatrix()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintMatrixSpiralOrder(tt.args.matrix)
		})
	}
}
