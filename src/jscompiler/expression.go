package jscompiler

import (
	"strings"

	"../parser"
)

const (
	corePrefix         = "__core"
	expressionExecuter = "E"
)

func handleExpressionName(name string) string {
	parts := strings.Split(name, "/")
	acc := ""

	for _, part := range parts {
		acc = acc + "[\"" + part + "\"]"
	}

	return acc
}

func generateWhitespace(l int) string {
	return strings.Repeat("  ", l)
}

func HandleExpression(ex *parser.Token, args string, depth int) string {
	acc := "\n" + generateWhitespace(depth) + expressionExecuter + "(" + CtxPrefix + handleExpressionName(ex.TokenValue) + ", " + args + ")"

	if ex.ParentToken.TokenType == "root" {
		acc = acc + ".call(" + CtxPrefix + ");"
	}

	return acc
}
