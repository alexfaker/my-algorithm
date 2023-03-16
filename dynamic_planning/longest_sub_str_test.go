package dynamic_planning

import "testing"

func Test_getLongestSubStr(t *testing.T) {
	type args struct {
		subStr1 string
		subStr2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{
				subStr1: "A1234",
				subStr2: "CD1234",
			},
			"1234",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLongestSubStr(tt.args.subStr1, tt.args.subStr2); got != tt.want {
				t.Errorf("getLongestSubStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLongestSubStrV2(t *testing.T) {
	type args struct {
		subStr1 string
		subStr2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{
				subStr1: "A1234",
				subStr2: "CD1234",
			},
			"1234",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLongestSubStrV2(tt.args.subStr1, tt.args.subStr2); got != tt.want {
				t.Errorf("getLongestSubStrV2() = %v, want %v", got, tt.want)
			}
		})
	}
}