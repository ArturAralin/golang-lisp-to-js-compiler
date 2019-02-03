package parser

import (
	"fmt"
	"logger"
	"os"
	"regexp"
	"stringstack"
)

const (
	namespace = "parser"
)

var fname string
var callDepth int

type Token struct {
	tokenValue  string
	tokenType   string
	parentToken *Token
	childTokens []*Token
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
	logger.Log(fname, namespace, "[START] parseFunctionName", callDepth)
	currentSymbol := ""
	fnName := ""

	for currentSymbol != " " {
		currentSymbol = string(input[pos])
		logger.UpdateLineAndColumn(currentSymbol)
		fnName = fnName + currentSymbol
		pos = pos + 1
	}

	logger.Log(fname, namespace, "[START] parseFunctionName", callDepth)

	return pos, fnName
}

func matchString(mask string, s string) bool {
	r, _ := regexp.MatchString(mask, s)

	return r
}

func getValueType(firstCharacter string) string {
	switch {
	case matchString("[\\.1234567890]", firstCharacter):
		return "number"
	case matchString("\"", firstCharacter):
		return "string"
	default:
		logger.ThrowError("Unknown type")
	}

	return ""
}

// TODO: add type matching
func parseValue(cursorPosition int, input string) (int, *Token) {
	// fmt.Println("[START] parseValue cursorPosition=", cursorPosition)
	currentSymbol := ""
	token := &Token{}

	token.tokenType = getValueType(string(input[cursorPosition]))

	for true {
		currentSymbol = string(input[cursorPosition])
		logger.UpdateLineAndColumn(currentSymbol)

		if isSpecSymbol(currentSymbol) {
			break
		}

		token.tokenValue = token.tokenValue + currentSymbol
		cursorPosition = cursorPosition + 1
	}

	// fmt.Println("[FINISH] parseValue cursorPosition=", cursorPosition)
	return cursorPosition, token
}

func Parse(fileName, input string) *Token {
	fname = fileName
	s := make(stringstack.Stack, 0)
	inputLen := len(input)
	cursorPosition := 0
	currentSymbol := ""
	currentToken := &Token{tokenType: "root"}

	for cursorPosition < inputLen {
		currentSymbol = string(input[cursorPosition])
		logger.UpdateLineAndColumn(currentSymbol)

		if currentSymbol == " " || currentSymbol == "\n" || currentSymbol == "\r" {
			cursorPosition = cursorPosition + 1
			continue
		}

		// start new token
		if currentSymbol == "(" {
			// TODO: add checking previous symbol
			// for checking syntax error
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

		// fmt.Println(currentToken, currentSymbol)
		callDepth = s.Len()
	}

	if s.Len() > 0 {
		fmt.Println("Syntax error")
		os.Exit(1)
	}

	return currentToken
}
