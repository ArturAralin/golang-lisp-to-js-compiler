package parser

const (
	JSCallMask = "[\\.a-zA-Z0-9_]"
)

func ParseJSCall(t *Token, cursorPosition int, input string) int {
	acc := ""
	currentSymbol := ""

	for true {
		currentSymbol = string(input[cursorPosition])

		if !MatchString(JSCallMask, currentSymbol) {
			break
		}

		acc = acc + currentSymbol
		cursorPosition = cursorPosition + 1
	}

	t.TokenValue = acc[1:]

	return cursorPosition
}
