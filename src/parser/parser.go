package parser

import (
	"fmt"
	"os"
	"regexp"

	"../logger"
	"../stringstack"
)

const (
	namespace          = "parser"
	openSpecialSymbol  = "[{\\[\\(]"
	closeSpecialSymbol = "[}\\]\\)]"
	functionSymbols    = "[a-z-?!]"
)

var fname string
var callDepth int

type Token struct {
	tokenValue  string
	tokenType   string
	parentToken *Token
	ChildTokens []*Token
}

func isAllowedToFnName(s string) bool {
	r, _ := regexp.MatchString("([-a-z])", s)

	return r
}

func isSpecSymbol(s string) bool {
	r, _ := regexp.MatchString("[\\(\\)\\[\\]]", s)

	return r
}

func parseExpressionName(pos int, input string) (int, string) {
	logger.Log(fname, namespace, "[START] parseFunctionName", callDepth)
	currentSymbol := ""
	fnName := ""
	// escape first symbol
	pos = pos + 1

	// TODO: allow only \w+ \d+ ? ! -
	for true {
		currentSymbol = string(input[pos])
		pos = pos + 1
		logger.UpdateLineAndColumn(currentSymbol)

		if !MatchString(functionSymbols, currentSymbol) {
			break
		}

		fnName = fnName + currentSymbol
	}

	logger.Log(fname, namespace, "[START] parsed function name \""+fnName+"\"", callDepth)

	return pos, fnName
}

func getValueType(cursorPosition int, input string) string {
	firstCharacter := string(input[cursorPosition])

	switch {
	case MatchString(NumberMatchMask, firstCharacter):
		return "number"
	case MatchString(StringMatchMask, firstCharacter):
		return "string"
	case MatchString(ReservedWordsMatch, firstCharacter) && len(FindReservedWord(cursorPosition, input)) > 0:
		return "reservedWord"
	default:
		logger.ThrowError("Unknown value type for character \"" + firstCharacter + "\"")
	}

	return ""
}

func parseValue(cursorPosition int, input string) (int, *Token) {
	logger.Log(fname, namespace, "[START] parseValue", callDepth)

	token := &Token{tokenType: getValueType(cursorPosition, input)}

	switch token.tokenType {
	case "number":
		cursorPosition = ParseNumber(token, cursorPosition, input)
	case "string":
		cursorPosition = ParseSting(token, cursorPosition, input)
	case "reservedWord":
		cursorPosition = ParseReservedWords(token, cursorPosition, input)
	default:
		logger.ThrowError("Not found parser for type \"" + token.tokenType + "\"")
	}

	logger.Log(fname, namespace, "[FINISH] parsedValue="+token.tokenValue, callDepth)

	return cursorPosition, token
}

func getType(s string) string {
	switch s {
	case "(":
		return "expression"
	case "[":
		return "array"
	case "{":
		return "object"
	default:
		logger.ThrowError("Unknown type")
	}

	return ""
}

func Parse(fileName, input string) *Token {
	fname = fileName
	s := make(stringstack.Stack, 0)
	inputLen := len(input)
	cursorPosition := 0
	currentSymbol := ""
	currentToken := &Token{tokenType: "root"}

	for cursorPosition < inputLen {
		callDepth = s.Len()
		currentSymbol = string(input[cursorPosition])
		logger.UpdateLineAndColumn(currentSymbol)

		if currentSymbol == " " || currentSymbol == "," || currentSymbol == "\n" || currentSymbol == "\r" {
			cursorPosition = cursorPosition + 1
			continue
		}

		// start new token
		if MatchString(openSpecialSymbol, currentSymbol) {
			// TODO: add checking previous symbol
			// for checking syntax error
			s = s.Push(currentSymbol)
			newToken := &Token{
				tokenType:   getType(currentSymbol),
				parentToken: currentToken,
			}

			// parse name
			if newToken.tokenType == "expression" {
				newSymbolPos, fnName := parseExpressionName(cursorPosition, input)
				currentToken.tokenValue = fnName
				cursorPosition = newSymbolPos
			}

			currentToken.ChildTokens = append(currentToken.ChildTokens, newToken)
			currentToken = newToken
			cursorPosition = cursorPosition + 1

			continue
		}

		// retrun to previous token
		if MatchString(closeSpecialSymbol, currentSymbol) {
			s, _ = s.Pop()

			// validate object
			if currentToken.tokenType == "object" {
				if (len(currentToken.ChildTokens) % 2) > 0 {
					logger.ThrowError("Object must have even elements")
				}
			}

			currentToken = currentToken.parentToken
			cursorPosition = cursorPosition + 1
			continue
		}

		newCursorPosition, t := parseValue(cursorPosition, input)
		cursorPosition = newCursorPosition
		t.parentToken = currentToken
		currentToken.ChildTokens = append(currentToken.ChildTokens, t)
	}

	if s.Len() > 0 {
		fmt.Println("Syntax error")
		os.Exit(1)
	}

	return currentToken
}
