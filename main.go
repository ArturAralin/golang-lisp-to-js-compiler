package main

import "fmt"
import "os"
import "regexp"

type Token struct {
	tokenValue  string
	tokenType   string
	parentToken *Token
	childTokens []*Token
}

type stack []string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, string) {
	l := len(s)

	if l == 0 {
		return s, ""
	}

	return s[:l-1], s[l-1]
}

func (s stack) Len() int {
	return len(s)
}

func isAllowedToFnName(s string) bool {
	r, _ := regexp.MatchString("([-a-z])", s)

	return r
}

func isSpecSymbol(s string) bool {
	r, _ := regexp.MatchString("[\\(\\)\\[\\] ]", s)

	return r
}

func parseFunctionName(pos int, input string) (int, string) {
	fmt.Println("[START] parseFunctionName cursorPosition=", pos)
	currentSymbol := ""
	fnName := ""

	for currentSymbol != " " {
		currentSymbol = string(input[pos])
		fnName = fnName + currentSymbol
		pos = pos + 1
	}

	fmt.Println("[FINISH] parseFunctionName cursorPosition=", pos)
	return pos, fnName
}

func parseValue(cursorPosition int, input string) (int, *Token) {
	fmt.Println("[START] parseValue cursorPosition=", cursorPosition)
	currentSymbol := ""
	token := &Token{}

	fmt.Println(currentSymbol)
	for true {
		currentSymbol = string(input[cursorPosition])

		if isSpecSymbol(currentSymbol) {
			break
		}

		token.tokenValue = token.tokenValue + currentSymbol
		cursorPosition = cursorPosition + 1
	}

	fmt.Println("[FINISH] parseValue cursorPosition=", cursorPosition)
	return cursorPosition, token
}

func main() {
	s := make(stack, 0)
	input := "(fn (sub-fn 1 2) 10)"
	inputLen := len(input)
	// column := 0
	// line := 0
	cursorPosition := 0
	currentSymbol := ""
	currentToken := &Token{tokenType: "root"}

	for cursorPosition < inputLen {
		currentSymbol = string(input[cursorPosition])

		if currentSymbol == " " {
			cursorPosition = cursorPosition + 1
			continue
		}

		// start new token
		if currentSymbol == "(" {
			s = s.Push("(")
			newToken := &Token{
				tokenType:   "function",
				parentToken: currentToken,
			}
			currentToken.childTokens = append(currentToken.childTokens, newToken)
			currentToken = newToken

			// parse tokenValue name
			newSymbolPos, fnName := parseFunctionName(cursorPosition+1, input)
			currentToken.tokenValue = fnName
			cursorPosition = newSymbolPos
			continue
		}

		// retrun to previous token
		if currentSymbol == ")" {
			s, _ = s.Pop()
			currentToken = currentToken.parentToken
			cursorPosition = cursorPosition + 1
			continue
		}

		newCursorPosition, t := parseValue(cursorPosition, input)
		cursorPosition = newCursorPosition

		t.parentToken = currentToken
		currentToken.childTokens = append(currentToken.childTokens, t)

		fmt.Println(currentToken, currentSymbol)
	}

	// fmt.Println(currentToken.childTokens[0].childTokens[0])

	if s.Len() > 0 {
		fmt.Println("Syntax error")
		os.Exit(1)
	}
}
