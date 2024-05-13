package pkgs

import (
	"strconv"
	"strings"
)

// IntsToString 将整型数组转换为逗号分隔的字符串
func IntsToString(ints []int64) string {
	var strArr []string
	for _, i := range ints {
		strArr = append(strArr, strconv.FormatInt(i, 10))
	}
	return strings.Join(strArr, ", ")
}
