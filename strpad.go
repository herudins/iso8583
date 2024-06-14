package iso8583

import "strings"

func leftPad(s string, length int, pad string) string {
	if len(s) >= length {
		return s
	}
	padding := strings.Repeat(pad, length-len(s))
	return padding + s
}

func rightPad(s string, l int, pad string) string {
	if len(s) >= l {
		return s
	}
	padding := strings.Repeat(pad, l-len(s))

	return s + padding
}
