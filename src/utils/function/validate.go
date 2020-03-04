package function

import (
	"errors"
	"regexp"
	"sort"
	"strings"
)

const (
	regularTel    = "^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(18[0,5-9]))\\d{8}$"
	regularCardNo = "^[1-9]\\d{7}((0\\d)|(1[0-2]))(([0|1|2]\\d)|3[0-1])\\d{3}$|^[1-9]\\d{5}[1-9]\\d{3}((0\\d)|(1[0-2]))(([0|1|2]\\d)|3[0-1])\\d{3}([0-9]|[xX])$"
)

//验证身份证号码
func ValidateCardNo(str string) bool {
	reg := regexp.MustCompile(regularCardNo)
	return reg.MatchString(str)
}

//验证手机号码
func ValidateTel(str string) bool {
	reg := regexp.MustCompile(regularTel)
	return reg.MatchString(str)
}

//验证非法字符
func VerifyInjection(str string) error {
	if strings.Index(str, "'") > -1 || strings.Index(str, "\"") > -1 || strings.Index(str, ";") > -1 {
		return errors.New("Your string has a special character")
	}
	return nil
}

//数组去重去空
func RemoveDuplicatesAndEmpty(a []string) (ret []string) {
	sort.Strings(a)
	a_len := len(a)
	for i := 0; i < a_len; i++ {
		if (i > 0 && a[i-1] == a[i]) || len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}
