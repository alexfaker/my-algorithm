package strings


// IsDeformation
// @Description: 判断两个字符串是否互为变形词
// @param str1
// @param str2
// @return bool
func IsDeformation(str1, str2 string) bool {
	var res = true
	var varCharMap1 = make(map[uint8]int)
	var varCharMap2 = make(map[uint8]int)
	if len(str1) == 0 || len(str2) == 0 {
		return false
	}
	for i := 0; i < len(str1); i++ {
		if _, ok := varCharMap1[str1[i]]; ok {
			varCharMap1[str1[i]] += 1
		} else {
			varCharMap1[str1[i]] = 1
		}
	}
	for i := 0; i < len(str2); i++ {
		if _, ok := varCharMap2[str2[i]]; ok {
			varCharMap2[str2[i]] += 1
		} else {
			varCharMap2[str2[i]] = 1
		}
	}
	for k, v := range varCharMap1 {
		if _, ok := varCharMap2[k]; !ok {
			res = false
			return res
		}
		if v != varCharMap2[k] {
			res = false
			return res
		}
	}
	return res
}
