package util

import "strings"

func SplitLines(str string) []string {
	lines := strings.Split(str, getDelim(str))
	return lines
}

func getDelim(str string) string {
	firstLineBreak := strings.Index(str, "\n")
	if firstLineBreak > 0 {
		if str[firstLineBreak-1] == '\r' {
			return "\r\n"
		}
		return "\n"
	}

	// just default to \n
	return "\n"
}
