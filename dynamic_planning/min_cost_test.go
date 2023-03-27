package dynamic_planning

import (
	"fmt"
	"testing"
)

func Test_generatePeople(t *testing.T) {
	type args struct {
		num    int
		groups [][]string
	}
	groupList := [][]string{
		{"小名", "小红", "小马", "小丽", "小强"},
		{"大壮", "大力", "大1", "大2", "大3"},
		{"阿花", "阿朵", "阿蓝", "阿紫", "阿红"},
		{"A", "B", "C", "D", "E"},
		{"一", "二", "三", "四", "五"},
		{"建国", "建军", "建民", "建超", "建跃"},
		{"爱民", "爱军", "爱国", "爱辉", "爱月"},
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"1",
			args{
				num:    7,
				groups: groupList,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateTeams(GeneratePeople(tt.args.num,tt.args.groups))
			for _, v := range got{
				fmt.Println(v)
			}
		})
	}
}
