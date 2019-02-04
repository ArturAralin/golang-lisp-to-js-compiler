package jscompiler

import (
	"fmt"

	"../parser"
)

func Compile(root *parser.Token) string {
	// currentToken := root
	var visited []*parser.Token
	var t []string
	var text = &t

	fmt.Println(text)

	// ?????
	// for visited[len(visited)-1] != root {
	// 	fmt.Println("")
	// }

	return ""
}

// expression -> function
