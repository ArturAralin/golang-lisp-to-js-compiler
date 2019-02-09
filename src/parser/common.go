package parser

import (
	"regexp"
)

func MatchString(mask string, s string) bool {
	r, _ := regexp.MatchString(mask, s)

	return r
}

// TODO: improve this file performance
func MatchByte(mask string, b byte) bool {
	l, err := regexp.MatchString(mask, string(b))

	if err != nil {
		return false
	}

	return l
}
