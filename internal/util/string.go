package util

import (
	"fmt"
	"reflect"
	"strings"
)

func JoinStrings(str ...any) string {
	var builder strings.Builder
	for i := range str {
		builder.WriteString(checkStr(str[i]))
	}
	return builder.String()
}

func checkStr(a any) string {
	if reflect.ValueOf(a).Kind() == reflect.String {
		return a.(string)
	} else {
		return fmt.Sprintf("%v", a)
	}
}
