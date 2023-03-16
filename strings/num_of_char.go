package strings

import (
	"regexp"
	"strconv"
	"strings"
)

func NumOfChar(str string) int {
	var ans int
	var le = len(str)
	if le == 0 {
		return ans
	}
	var r, _ = regexp.Compile("[-0-9]+")
	res := r.FindAllString(str, -1)
	for _, v := range res {
		var t int
		var isPositive bool
		v = strings.TrimRight(v,"-")
		for strings.Contains(v, "-") {
			v = v[1:]
			t += 1
		}
		if t%2 == 0 {
			isPositive = true
		}
		tmp, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		if isPositive {
			ans += tmp
		} else {
			ans += tmp * -1

		}
	}
	return ans
}
