package logger

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	whitespaceSymbol = "  "
)

var column = 1
var line = 1

func intToStr(v int) string {
	return strconv.Itoa(v)
}

func UpdateLineAndColumn(s string) {
	column = column + 1

	if s == "\n" {
		column = 0
		line = line + 1
	}
}

func ThrowError(m string) {
	fmt.Printf("%s:%s %s\n", intToStr(line), intToStr(column), m)
	os.Exit(1)
}

func Log(fileName, ns, s string, callDepth int) {
	whitespace := strings.Repeat(whitespaceSymbol, callDepth)
	file := fileName + ":" + intToStr(line) + ":" + intToStr(column)
	t := whitespace + " [" + ns + "]" + s + " " + file
	fmt.Println(t)
}
