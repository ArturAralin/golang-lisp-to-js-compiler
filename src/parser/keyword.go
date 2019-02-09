package parser

import "strings"

const (
	KeywordMask = "[a-zA-Z-?!/0-9]"
)

func ParseKeyword(t *Token, cursorPosition int, input string) int {
	var acc strings.Builder
	var currentSymbol byte

	// skip first character ":"
	cursorPosition++

	for true {
		currentSymbol = input[cursorPosition]

		if !MatchByte(KeywordMask, currentSymbol) {
			break
		}

		acc.WriteByte(currentSymbol)
		cursorPosition++
	}

	t.TokenValue = acc.String()
	return cursorPosition
}
