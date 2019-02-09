package parser

const (
	SymbolSymbols = "[a-zA-Z-?!/0-9]"
)

func ParseSymbol(token *Token, cursorPosition int, input string) int {
	acc := ""
	currentSymbol := ""

	for true {
		currentSymbol = string(input[cursorPosition])

		if !MatchString(SymbolSymbols, currentSymbol) {
			break
		}

		acc = acc + currentSymbol
		cursorPosition = cursorPosition + 1
	}

	token.TokenValue = acc

	return cursorPosition
}
