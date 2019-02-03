package parser

import (
	"regexp"
)

func MatchString(mask string, s string) bool {
	r, _ := regexp.MatchString(mask, s)

	return r
}
