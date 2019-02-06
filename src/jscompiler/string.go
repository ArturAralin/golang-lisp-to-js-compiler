package jscompiler

import "../parser"

func HandleString(node *parser.Token) string {
	return "\"" + node.TokenValue + "\""
}
