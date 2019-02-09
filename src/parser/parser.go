package parser

import (
	"fmt"
	"os"
	"regexp"

	"../logger"
	"../stringstack"
)

const (
	namespace             = "parser"
	openSpecialSymbol     = "[{\\[\\(]"
	closeSpecialSymbol    = "[}\\]\\)]"
	contextDetachedSymbol = "'"
)

var fname string
var callDepth int

type Token struct {
	TokenValue      string
	TokenType       string
	ParentToken     *Token
	ChildTokens     []*Token
	ContextDetached bool
}

func isSpecSymbol(s string) bool {
	r, _ := regexp.MatchString("[\\(\\)\\[\\]]", s)

	return r
}

func parseExpressionName(pos int, input string) (int, string) {
	logger.Log(fname, namespace, "[START] parseFunctionName", callDepth)
	currentSymbol := ""
	fnName := ""
	// escape first symbol (
	pos = pos + 1

	for true {
		currentSymbol = string(input[pos])
		pos = pos + 1
		logger.UpdateLineAndColumn(currentSymbol)

		if !MatchString(SymbolSymbols, currentSymbol) {
			break
		}

		fnName = fnName + currentSymbol
	}

	logger.Log(fname, namespace, "[START] parsed function name \""+fnName+"\"", callDepth)

	return pos - 1, fnName
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
	case MatchString(SymbolSymbols, firstCharacter):
		return "symbol"
	default:
		logger.ThrowError("Unknown value type for character \"" + firstCharacter + "\"")
	}

	return ""
}

func parseValue(cursorPosition int, input string) (int, *Token) {
	logger.Log(fname, namespace, "[START] parseValue", callDepth)

	token := &Token{TokenType: getValueType(cursorPosition, input)}

	switch token.TokenType {
	case "number":
		cursorPosition = ParseNumber(token, cursorPosition, input)
	case "string":
		cursorPosition = ParseSting(token, cursorPosition, input)
	case "reservedWord":
		cursorPosition = ParseReservedWords(token, cursorPosition, input)
	case "symbol":
		cursorPosition = ParseSymbol(token, cursorPosition, input)
	default:
		logger.ThrowError("Not found parser for type \"" + token.TokenType + "\"")
	}

	logger.Log(fname, namespace, "[FINISH] parsedValue="+token.TokenValue, callDepth)

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
		return "symbol"
	}
}

func Parse(fileName, input string) *Token {
	fname = fileName
	s := make(stringstack.Stack, 0)
	inputLen := len(input)
	cursorPosition := 0
	currentSymbol := ""
	prevSymbol := ""
	currentToken := &Token{TokenType: "root"}
	root := currentToken

	for cursorPosition < inputLen {
		callDepth = s.Len()
		prevSymbol = currentSymbol
		currentSymbol = string(input[cursorPosition])
		logger.UpdateLineAndColumn(currentSymbol)

		if currentSymbol == " " || currentSymbol == "," || currentSymbol == "\n" || currentSymbol == "\r" {
			cursorPosition = cursorPosition + 1
			continue
		}

		// handle ctx detach symbol
		if currentSymbol == contextDetachedSymbol {
			cursorPosition = cursorPosition + 1

			if !MatchString(openSpecialSymbol, string(input[cursorPosition])) {
				logger.ThrowError("Unexpected context detach symbol")
			}

			continue
		}

		// start new token
		if MatchString(openSpecialSymbol, currentSymbol) {
			// TODO: add checking previous symbol
			// for checking syntax error
			s = s.Push(currentSymbol)

			newToken := &Token{
				TokenType:       getType(currentSymbol),
				ParentToken:     currentToken,
				ContextDetached: currentToken.ContextDetached || prevSymbol == contextDetachedSymbol,
			}

			// parse name
			if newToken.TokenType == "expression" {
				newSymbolPos, fnName := parseExpressionName(cursorPosition, input)
				newToken.TokenValue = fnName
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
			if currentToken.TokenType == "object" {
				if (len(currentToken.ChildTokens) % 2) > 0 {
					logger.ThrowError("Object must have even elements")
				}
			}

			currentToken = currentToken.ParentToken
			cursorPosition = cursorPosition + 1
			continue
		}

		newCursorPosition, t := parseValue(cursorPosition, input)
		cursorPosition = newCursorPosition
		t.ParentToken = currentToken
		currentToken.ChildTokens = append(currentToken.ChildTokens, t)
	}

	if s.Len() > 0 {
		fmt.Println("Syntax error")
		os.Exit(1)
	}

	return root
}
