package function

import (
	"strings"
)

func MoveXing(str string) (result string) {
	strArr := strings.Split(str, "")
	right := len(strArr) - 1
	left := right

	for left > 0 {
		if strArr[right] != "*" {
			right--
			if left > right {
				left = right
			}
		} else {
			for left > 0 && strArr[left] == "*" {
				left--
			}

			if left > 0 {
				strArr[right] = strArr[left]
				strArr[left] = "*"
				left--
				right--
			}
		}
	}

	for _, item := range strArr {
		result += item
	}

	return result
}

func MoveXing2(str string) (result string) {
	for _, item := range str {
		if item == '*' {
			result += "*"
		}
	}

	for _, item := range str {
		if item != '*' {
			result += string(item)
		}
	}
	return result
}

func MoveXing3(str string) string {
	str1 := ""
	str2 := ""

	for _, item := range str {
		if item == '*' {
			str1 += "*"
		} else {
			str2 += string(item)
		}
	}

	return str1 + str2
}
