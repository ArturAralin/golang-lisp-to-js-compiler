package parser

const (
	StringMatchMask = "\""
	escapeSymbol    = "\\"
	endStringSymbol = "\""
)

func ParseString(token *Token, cursorPosition int, input string) int {
	currentSymbol := ""
	l := 0
	isHaveEscapeSymbol := false
	acc := ""

	// skip first quote
	cursorPosition = cursorPosition + 1

	for true {
		currentSymbol = string(input[cursorPosition])

		if currentSymbol == escapeSymbol {
			isHaveEscapeSymbol = true
		}

		if currentSymbol == endStringSymbol && isHaveEscapeSymbol == false {
			break
		}

		if currentSymbol == endStringSymbol && isHaveEscapeSymbol == true {
			isHaveEscapeSymbol = false
		}

		acc = acc + currentSymbol

		l = l + 1
		cursorPosition = cursorPosition + 1
	}

	token.TokenValue = acc

	return cursorPosition + 1
}
