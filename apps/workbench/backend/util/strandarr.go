package util

import "strings"

// TransStrToArr Convert a string to an array of strings
func TransStrToArr(portStr string) []string {
	portStr = strings.Replace(portStr, "[", "", 1)
	portStr = strings.Replace(portStr, "]", "", 1)
	portArr := strings.Split(portStr, ",")
	return portArr
}

// TransArrToStr String array converted to string
func TransArrToStr(portArr []string) string {
	str := strings.Join(portArr, ",")
	str = "[" + str + "]"
	return str
}
