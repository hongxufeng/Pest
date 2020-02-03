package function

import (
	"errors"
	"sort"
	"strings"
)

func VerifyInjection(str string) error {
	if strings.Index(str, "'") > -1 || strings.Index(str, "\"") > -1 || strings.Index(str, ";") > -1 {
		return errors.New("Your string has a special character")
	}
	return nil
}
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
