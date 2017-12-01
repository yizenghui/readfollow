package repository

import (
	"strconv"
	"strings"
)

//Str2Int64 url请求中使用的字符串转int64数组
func Str2Int64(str string) []int64 {
	var tags []int64
	arr := strings.Split(str, ",")
	for _, v := range arr {
		val, err := strconv.Atoi(v)
		if err == nil && val > 0 {
			tags = append(tags, int64(val))
		}
	}
	return tags
}
