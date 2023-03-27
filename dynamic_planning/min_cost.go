package dynamic_planning
import (
	"fmt"
	"math/rand"
	"time"
)

// 计算文本最小编辑距离问题
/*
给定 字符串 str1, str2 ,再给定三个整数 ic, dc, rc 分别代表
插入,删除, 替换一个字符串的代价,返回将 str1 修改为 str2的最小代价
*/

// 定义一个结构体表示一个人
type Person struct {
	Name  string // 名字
	Group int    // 所在组
}

// 定义一个函数，用于生成指定数量的人员
func GeneratePeople(num int, groups [][]string) []Person {
	people := make([]Person, 0)
	for _, group := range groups {
		for _, name := range group {
			people = append(people, Person{Name: name, Group: len(people) / (num / len(groups))})
		}
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(people), func(i, j int) {
		people[i], people[j] = people[j], people[i]
	})
	return people
}

// 定义一个函数，用于打印队伍
func printTeams(teams [][]Person) {
	fmt.Println("Teams:")
	for _, team := range teams {
		for _, person := range team {
			fmt.Printf("%s ", person.Name)
		}
		fmt.Println()
	}
}

// 定义一个函数，用于检查是否所有队伍都满足要求
func checkTeams(teams [][]Person) bool {
	if len(teams) < 2 {
		return false
	}
	for _, team := range teams {
		if len(team) < 2 || len(team) > 3 {
			return false
		}
		if team[0].Group == team[1].Group {
			return false
		}
		if len(team) == 3 && team[1].Group == team[2].Group {
			return false
		}
	}
	return true
}

// 定义一个函数，用于生成随机的队伍
func GenerateTeams(people []Person) [][]Person {
	for {
		teams := make([][]Person, 0)
		used := make([]bool, len(people))
		for i := 0; i < len(people); {
			if used[i] {
				i++
				continue
			}
			teamSize := 2
			if rand.Float32() < 0.5 {
				teamSize = 3
			}
			team := make([]Person, teamSize)
			groupUsed := make([]bool, len(people)/teamSize)
			for j := 0; j < teamSize; {
				if j == 1 && teamSize == 3 {
					groupUsed[people[i].Group] = true
				}
				if groupUsed[people[i].Group] || used[i] {
					i++
					continue
				}
				team[j] = people[i]
				groupUsed[people[i].Group] = true
				used[i] = true
				j++
				i++
			}
			teams = append(teams, team)
		}
		if checkTeams(teams) {
			return teams
		}
	}
}
