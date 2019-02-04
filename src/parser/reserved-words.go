package parser

const (
	ReservedWordsMatch = "[ntf]"
)

func checkWord(word string, cursorPosition int, input string) bool {
	for _, char := range word {
		if string(input[cursorPosition]) != string(char) {
			return false
		}

		cursorPosition = cursorPosition + 1
	}

	return true
}

func FindReservedWord(cursorPosition int, input string) string {
	firstCharacter := string(input[cursorPosition])

	if firstCharacter == "t" {
		if checkWord("true", cursorPosition, input) {
			return "true"
		}
	}

	if firstCharacter == "f" {
		if checkWord("false", cursorPosition, input) {
			return "false"
		}
	}

	if firstCharacter == "n" {
		if checkWord("nil", cursorPosition, input) {
			return "nil"
		}
	}

	return ""
}

func ParseReservedWords(token *Token, cursorPosition int, input string) int {
	word := FindReservedWord(cursorPosition, input)

	token.TokenValue = word

	return cursorPosition + len(word)
}

// n nil
// t true
// f false
