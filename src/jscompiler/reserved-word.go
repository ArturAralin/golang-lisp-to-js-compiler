package jscompiler

import (
	"../parser"
)

func HandleReservedWord(t *parser.Token) string {
	if t.TokenValue == "nil" {
		return "null"
	}

	return t.TokenValue
}
