package jscompiler

import (
	"strings"

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
			case "reservedWord":
				acc = acc + HandleReservedWord(chT)
			case "expression":
				acc = acc + HandleExpression(chT, walkTree(chT, depth+1), depth)
			case "number":
				acc = acc + chT.TokenValue
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
			case "keyword":
				acc = acc + "keyword('" + chT.TokenValue + "')"
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
	var js strings.Builder

	js.WriteString("var __core = require(\"fjs-compiler/lib/core/core.js\");\n")
	js.WriteString("var ")
	js.WriteString(CtxPrefix)
	js.WriteString(" = Object.assign({}, __core);\n")
	js.WriteString(CtxPrefix)
	js.WriteString(".ROOT = ")
	js.WriteString(CtxPrefix)
	js.WriteString(";\n")
	js.WriteString("var symbol = __core.$;\n")
	js.WriteString("var E = __core[symbol('E')];\n")
	js.WriteString("var keyword = __core[symbol('keyword')];\n")
	js.WriteString("\n")

	code := walkTree(root, 0)

	js.WriteString(code)

	if isTestMode {
		js.WriteString("\n\n// this export added by compiler\n")
		js.WriteString("exports.CTX = ")
		js.WriteString(CtxPrefix)
		js.WriteString(";\n")
	}

	return js.String()
}
