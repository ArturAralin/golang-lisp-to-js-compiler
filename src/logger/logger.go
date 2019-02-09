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

var logEnabled = len(os.Getenv("LOG_ENABLED")) > 0

var Column = 1
var line = 1

func intToStr(v int) string {
	return strconv.Itoa(v)
}

func UpdateLineAndColumn(s string) {
	Column = Column + 1

	if s == "\n" {
		Column = 0
		line = line + 1
	}
}

func ThrowError(m string) {
	fmt.Printf("%s:%s %s\n", intToStr(line), intToStr(Column), m)
	os.Exit(1)
}

func Log(fileName, ns, s string, callDepth int) {
	if logEnabled {
		whitespace := strings.Repeat(whitespaceSymbol, callDepth)
		t := whitespace + "// [" + ns + "]" + s
		fmt.Println(t)
	}
}
