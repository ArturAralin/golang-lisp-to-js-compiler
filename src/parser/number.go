package parser

import (
	"../logger"
)

const (
	NumberMatchMask = "[\\.1234567890\\-+e]"
	exponentSymbol  = "e"
)

func ParseNumber(token *Token, cursorPosition int, input string) int {
	acc := ""
	currentSymbol := ""
	isHaveExponentSymbol := false
	l := 0

	for true {
		currentSymbol = string(input[cursorPosition])

		if currentSymbol == exponentSymbol && l == 0 {
			logger.ThrowError("Invalid number. Number must start by digit")
		}

		if currentSymbol == exponentSymbol && isHaveExponentSymbol == true {
			logger.ThrowError("Invalid number. Multiple \"e\" symbol")
		}

		if currentSymbol == exponentSymbol {
			isHaveExponentSymbol = true
		}

		if !MatchString(NumberMatchMask, currentSymbol) && string(acc[len(acc)-1]) == exponentSymbol {
			logger.ThrowError("Invalid number. After \"e\" must be a digit")
		}

		if !MatchString(NumberMatchMask, currentSymbol) {
			break
		}

		acc = acc + currentSymbol

		logger.UpdateLineAndColumn(currentSymbol)
		cursorPosition = cursorPosition + 1
		l = l + 1
	}

	token.TokenValue = acc
	return cursorPosition
}
