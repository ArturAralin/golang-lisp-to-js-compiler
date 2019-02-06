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
			switch chT.TokenType {
			case "expression":
				acc = acc + HandleExpression(chT, walkTree(chT))
			case "number":
				acc = acc + chT.TokenValue
			case "string":
				acc = acc + HandleString(chT)
			case "object":
				acc = acc + "{" + walkTree(chT) + "}"
			case "array":
				acc = acc + "[" + walkTree(chT) + "]"
			}

			if chT.ParentToken.TokenType == "object" {
				if idx%2 == 0 {
					acc = acc + ": "
				} else {
					acc = acc + ","
				}
			}

			if chT.ParentToken.TokenType == "array" {
				acc = acc + ", "
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
