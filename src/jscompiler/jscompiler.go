package jscompiler

import (
	"fmt"

	"../parser"
)

const (
	CtxPrefix  = "F"
	isTestMode = true
)

func walkTree(node *parser.Token, depth int) string {
	acc := ""
	tokensLen := len(node.ChildTokens)
	if tokensLen > 0 {
		for idx, chT := range node.ChildTokens {
			if chT.TokenType != "jsCall" &&
				chT.ParentToken.TokenType == "expression" &&
				idx == 0 {
				acc = acc + "F["
			}

			switch chT.TokenType {
			case "expression":
				acc = acc + HandleExpression(chT, walkTree(chT, depth+1), depth)
			case "number":
			case "jsCall":
				acc = acc + chT.TokenValue
			case "string":
				acc = acc + HandleString(chT)
			case "object":
				acc = acc + "{" + walkTree(chT, depth+1) + "}"
			case "array":
				acc = acc + "[" + walkTree(chT, depth+1) + "]"
			case "symbol":
				acc = acc + "symbol('" + chT.TokenValue + "')"
			}

			if chT.TokenType != "jsCall" &&
				chT.ParentToken.TokenType == "expression" &&
				idx == 0 {
				acc = acc + "]"
			}

			if chT.ParentToken.TokenType == "object" {
				if idx%2 == 0 {
					acc = acc + ": "
				} else {
					acc = acc + ","
				}
			}

			if (chT.ParentToken.TokenType == "expression" ||
				chT.ParentToken.TokenType == "array") && tokensLen-1 > idx {
				acc = acc + ", "
			}
		}
	}

	return acc
}

func Compile(root *parser.Token) string {
	header := "const __core = require(\"fjs-compiler/lib/core/core.js\");\n" +
		"const " + CtxPrefix + " = {...__core};\n" +
		CtxPrefix + ".ROOT = " + CtxPrefix + ";\n" +
		"const symbol = __core.$;\n" +
		"const E = __core[symbol('E')];\n"

	if isTestMode {
		header = header + "exports.CTX = " + CtxPrefix + ";\n"
	}

	header = header + "\n"

	code := walkTree(root, 0)

	fmt.Println(header + code)

	return ""
}

// expression -> function
