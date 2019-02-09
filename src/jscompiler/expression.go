package jscompiler

import (
	"strings"

	"../parser"
)

const (
	corePrefix         = "__core"
	expressionExecuter = "E"
)

func generateWhitespace(l int) string {
	return strings.Repeat("  ", l)
}

func HandleExpression(t *parser.Token, args string, depth int) string {
	acc := "\n" + generateWhitespace(depth) + expressionExecuter + "(" + args + ")"

	if t.ParentToken.TokenType == "root" {
		acc = acc + ".call(" + CtxPrefix + ");"
	}

	return acc
}
