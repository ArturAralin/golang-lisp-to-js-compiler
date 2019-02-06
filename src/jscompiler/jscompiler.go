package jscompiler

import (
	"fmt"

	"../parser"
)

const (
	CtxPrefix  = "F"
	isTestMode = true
)

func walkTree(node *parser.Token) string {
	acc := ""
	tokensLen := len(node.ChildTokens)
	if tokensLen > 0 {
		for idx, chT := range node.ChildTokens {
			if chT.TokenType == "expression" {
				acc = acc + HandleExpression(chT, walkTree(chT))
			}

			if chT.TokenType == "number" {
				acc = acc + chT.TokenValue
			}

			if chT.TokenType == "string" {
				acc = acc + "\"" + chT.TokenValue + "\""
			}

			if chT.TokenType == "expression" && chT.ParentToken.TokenType == "root" {
				acc = acc + ";\n"
			}

			if chT.ParentToken.TokenType == "expression" && tokensLen-1 > idx {
				acc = acc + ", "
			}
		}
	}

	return acc
}

func Compile(root *parser.Token) string {
	header := "const __core = require(\"fjs-compiler/lib/core/core.js\");\n" +
		"const " + CtxPrefix + " = {...__core};\n" +
		"const E = __core.E;\n"

	if isTestMode {
		header = header + "exports.CTX = " + CtxPrefix + ";\n"
	}

	header = header + "\n"

	code := walkTree(root)

	fmt.Println(header + code)

	return ""
}

// expression -> function
