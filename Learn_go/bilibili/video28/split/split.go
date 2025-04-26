package split

import "strings"

// split.go

// Split 将s按照sep进行分割，返回一个字符串的切片
// Split("香格里拉·弗陇提亚", "香格里拉") => ["·","弗陇提亚"]
func Split(s, sep string) (ret []string) {
	idx := strings.Index(s, sep)
	for idx > -1 {
		ret = append(ret, s[:idx])
		s = s[idx+1:]
		idx = strings.Index(s, sep)
	}
	ret = append(ret, s)
	return
}
