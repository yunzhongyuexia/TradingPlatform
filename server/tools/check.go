package tools

import (
	"regexp"
)

// CheckPhone 验证手机号码是否符合中国大陆的手机号码格式
func CheckPhone(phone string) bool {
	// 正则表达式匹配中国大陆手机号码
	// 1开头，第二位3-9，后面跟随9位数字
	reg := regexp.MustCompile(`^1[3-9]\d{9}$`)
	return reg.MatchString(phone)
}