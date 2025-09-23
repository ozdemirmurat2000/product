package utils

import "strings"

func Capitalize(str string) string {
	if len(str) == 0 {
		return str
	}
	// İlk harfi al → strings.ToUpper
	first := strings.ToUpper(string(str[0]))
	// Geri kalanını al → strings.ToLower
	rest := strings.ToLower(str[1:])
	return first + rest
}

func CapitalizeAllSmall(str string) string {
	return strings.ToLower(str)
}

func CapitaliseAllUpper(str string) string {
	return strings.ToUpper(str)
}
