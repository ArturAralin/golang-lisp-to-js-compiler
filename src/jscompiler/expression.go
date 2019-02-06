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

	return CtxPrefix + acc
}

func HandleExpression(ex *parser.Token, args string) string {
	acc := expressionExecuter + "(" + CtxPrefix + "." + ex.TokenValue + ", " + args + ")"

	if ex.ParentToken.TokenType == "root" {
		acc = acc + ".call(" + CtxPrefix + ");"
	}

	return acc
}
