package rule

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
)

var xvalues = [256]byte{
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 255, 255, 255, 255, 255, 255,
	255, 10, 11, 12, 13, 14, 15, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 10, 11, 12, 13, 14, 15, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
}

// 不能为空
func required(value string) bool {
	if value == "" {
		return false
	}
	return true
}

// 电邮地址
func email(value string) bool {
	rule := `^[a-z0-9]+([\-_\.][a-z0-9]+)*@([a-z0-9]+(-[a-z0-9]+)*\.)+[a-z]{2,4}$`
	return regexp.MustCompile(rule).MatchString(value)
}

// ip地址
func ip(str string) bool {
	rule := "^((2[0-4]\\d|25[0-5]|[01]?\\d\\d?)\\.){3}(2[0-4]\\d|25[0-5]|[01]?\\d\\d?)$"
	return regexp.MustCompile(rule).MatchString(str)
}

// 小写字母
func lower(str string) bool {
	rule := "^[a-z]*$"
	return regexp.MustCompile(rule).MatchString(str)
}

// 大写字母
func upper(str string) bool {
	rule := `^[a-z0-9A-Z\p{Han}]+(_[a-z0-9A-Z\p{Han}]+)*$`
	return regexp.MustCompile(rule).MatchString(str)
}

// 大小写字母
func letter(str string) bool {
	rule := `^[a-zA-Z]*$`
	return regexp.MustCompile(rule).MatchString(str)
}

// 整数(含0)
func intRule(str string) bool {
	rule := `^(-?[1-9]\d*|0)$`
	return regexp.MustCompile(rule).MatchString(str)
}

// 正整数(含0)
func uintRule(str string) bool {
	rule := `^([1-9]\d*|0)$`
	return regexp.MustCompile(rule).MatchString(str)
}

// 负整数
func nint(str string) bool {
	rule := `^-[1-9]\d*$`
	return regexp.MustCompile(rule).MatchString(str)
}

// 小数(含0)
func floatRule(str string) bool {
	if str == "0" {
		return true
	}
	rule := `^(?:[1-9][0-9]*(?:\.[0-9]+)?|0\.[0-9]+)$`
	return regexp.MustCompile(rule).MatchString(str)
}

// 正小数(含0)
func pfloat(str string) bool {
	rule := `^\d+(\.\d+)?$`
	return regexp.MustCompile(rule).MatchString(str)
}

// 负小数
func nfloat(str string) bool {
	rule := `^-([1-9]\d*\.\d*|0\.\d*[1-9]\d*)$`
	return regexp.MustCompile(rule).MatchString(str)
}

// 小写字母和数字
func lowerAndDigit(str string) bool {
	rule := `^[a-z0-9]*$`
	return regexp.MustCompile(rule).MatchString(str)
}

// 大写字母和数字
func upperAndDigit(str string) bool {
	rule := `^[A-Z0-9]*$`
	return regexp.MustCompile(rule).MatchString(str)
}

// 字母和数字
func letterAndDigit(str string) bool {
	rule := `^[a-zA-Z0-9]*$`
	return regexp.MustCompile(rule).MatchString(str)
}

// 中国大陆地区座机号码
func chineseTel(str string) bool {
	rule := `^(\(\d{3,4}\)|\d{3,4}-)?\d{7,8}$`
	return regexp.MustCompile(rule).MatchString(str)
}

// 中国大陆地区手机号码
func chineseMobile(str string) bool {
	rule := `^1(3|4|5|7|8)\d{9}$`
	return regexp.MustCompile(rule).MatchString(str)
}

// 中文
func chinese(str string) bool {
	rule := `^[\p{Han}]+$`
	return regexp.MustCompile(rule).MatchString(str)
}

// JSON类型
func jsonRule(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

func xtob(x1, x2 byte) (byte, bool) {
	b1 := xvalues[x1]
	b2 := xvalues[x2]
	return (b1 << 4) | b2, b1 != 255 && b2 != 255
}
func uuidRule(str string) bool {
	var uuid [16]byte
	switch len(str) {
	// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	case 36:

	// urn:uuid:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	case 36 + 9:
		if strings.ToLower(str[:9]) != "urn:uuid:" {
			return false
		}
		str = str[9:]

	// {xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}
	case 36 + 2:
		str = str[1:]

	// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	case 32:
		var ok bool
		for i := range uuid {
			uuid[i], ok = xtob(str[i*2], str[i*2+1])
			if !ok {
				return false
			}
		}
		return true
	default:
		return false
	}
	// s is now at least 36 bytes long
	// it must be of the form  xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	if str[8] != '-' || str[13] != '-' || str[18] != '-' || str[23] != '-' {
		return false
	}
	for i, x := range [16]int{
		0, 2, 4, 6,
		9, 11,
		14, 16,
		19, 21,
		24, 26, 28, 30, 32, 34} {
		v, ok := xtob(str[x], str[x+1])
		if !ok {
			return false
		}
		uuid[i] = v
	}
	return true
}

// 中国大陆地区身份证号码
func chineseIdentityCard(str string) bool {
	var idV int
	if str[17:] == "X" {
		idV = 88
	} else {
		var err error
		if idV, err = strconv.Atoi(str[17:]); err != nil {
			return false
		}
	}

	var verify int
	id := str[:17]
	arr := make([]int, 17)
	for i := 0; i < 17; i++ {
		arr[i], _ = strconv.Atoi(string(id[i]))
	}
	wi := [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	var res int
	for i := 0; i < 17; i++ {
		res += arr[i] * wi[i]
	}
	verify = res % 11

	var temp int
	a18 := [11]int{1, 0, 88 /* 'X' */, 9, 8, 7, 6, 5, 4, 3, 2}
	for i := 0; i < 11; i++ {
		if i == verify {
			temp = a18[i]
			break
		}
	}
	if temp == idV {
		return true
	}

	return false
}
