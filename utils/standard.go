package utils

import (
	"fmt"
	"unicode"
)

//驼峰 转 下划线
func Camel2Underline(camel string) (underline string) {
	for i, r := range camel {
		if unicode.IsUpper(r) {
			if i != 0 {
				underline = underline + "_"
			}
			underline = fmt.Sprintf("%s%c", underline, unicode.ToLower(r))
		} else {
			underline = fmt.Sprintf("%s%c", underline, unicode.ToLower(r))
		}
	}
	return
}
