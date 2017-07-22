package common

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/yizenghui/sda/code"
)

// TransformBookTotal 数据转换
func TransformBookTotal(total string) int64 {
	num := `^(?P<f>\d+)$`
	if b, _ := regexp.MatchString(num, total); b {
		t := code.FindString(num, total, "f")
		// 字符串转int64
		v, _ := strconv.ParseInt(t, 10, 64)
		return v
	}
	fl := `^(?P<f>\d+).(?P<l>\d+)$`
	if b, _ := regexp.MatchString(fl, total); b {
		Map := code.SelectString(fl, total)
		s := fmt.Sprintf("%v%v00", Map["f"], Map["l"])
		v, _ := strconv.ParseInt(s, 10, 64)
		return v
	}

	flw := `^(?P<f>\d+)万$`
	if b, _ := regexp.MatchString(flw, total); b {
		Map := code.SelectString(flw, total)
		s := fmt.Sprintf("%v0000", Map["f"])
		v, _ := strconv.ParseInt(s, 10, 64)
		return v
	}

	flb := `^(?P<f>\d+).(?P<l>\d+)万$`
	if b, _ := regexp.MatchString(flb, total); b {
		Map := code.SelectString(flb, total)
		s := fmt.Sprintf("%v%v00", Map["f"], Map["l"])
		v, _ := strconv.ParseInt(s, 10, 64)
		return v
	}
	return 0
}

//PrintBookTotal 显示字数
func PrintBookTotal(total int64) string {
	if total < 10000 {
		return strconv.FormatInt(total, 10)
	}
	t := float64(total) / float64(10000)
	v := strconv.FormatFloat(t, 'f', 2, 64)
	return fmt.Sprintf("%v万", v)
}
