package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"stringstack"
)

var column = 1
var line = 1

func intToStr(v int) string {
	return strconv.Itoa(v)
}
func throwError(m string) {
	fmt.Printf("%s:%s %s\n", intToStr(line), intToStr(column), m)
	os.Exit(1)
}

func updateLineAndColumn(s string) {
	column = column + 1

	if s == "\n" {
		column = 0
		line = line + 1
	}
}

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
	fmt.Println("[START] parseFunctionName cursorPosition=", pos)
	currentSymbol := ""
	fnName := ""

	for currentSymbol != " " {
		currentSymbol = string(input[pos])
		updateLineAndColumn(currentSymbol)
		fnName = fnName + currentSymbol
		pos = pos + 1
	}

	fmt.Println("[FINISH] parseFunctionName cursorPosition=", pos)
	return pos, fnName
}

func matchString(mask string, s string) bool {
	r, _ := regexp.MatchString(mask, s)

	return r
}

func getValueType(firstCharacter string) string {
	fmt.Println("firts char", firstCharacter)
	switch {
	case matchString("[\\.1234567890]", firstCharacter):
		return "number"
	case matchString("\"", firstCharacter):
		return "string"
	default:
		throwError("Unknown type")
	}

	return ""
}

// TODO: add type matching
func parseValue(cursorPosition int, input string) (int, *Token) {
	fmt.Println("[START] parseValue cursorPosition=", cursorPosition)
	currentSymbol := ""
	token := &Token{}

	token.tokenType = getValueType(string(input[cursorPosition]))

	for true {
		currentSymbol = string(input[cursorPosition])
		updateLineAndColumn(currentSymbol)

		if isSpecSymbol(currentSymbol) {
			break
		}

		token.tokenValue = token.tokenValue + currentSymbol
		cursorPosition = cursorPosition + 1
	}

	fmt.Println("[FINISH] parseValue cursorPosition=", cursorPosition)
	return cursorPosition, token
}

func readFile() string {
	d, _ := ioutil.ReadFile("./code")

	return string(d)
}

func main() {
	s := make(stringstack.Stack, 0)
	input := readFile()
	inputLen := len(input)
	cursorPosition := 0
	currentSymbol := ""
	currentToken := &Token{tokenType: "root"}

	for cursorPosition < inputLen {
		currentSymbol = string(input[cursorPosition])
		updateLineAndColumn(currentSymbol)

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

		fmt.Println(currentToken, currentSymbol)
	}

	if s.Len() > 0 {
		fmt.Println("Syntax error")
		os.Exit(1)
	}
}
