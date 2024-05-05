package main

import (
	"fmt"
	"strings"
)

func main() {
	decipherAndPrint("eqfg au hwp", "code")
	decipherAndPrint("eqfg# ku hwp!", "code")
	decipherAndPrint("eqfg ku hwp", "hamada")
}

func decipherAndPrint(cipheredMessage string, substring string) {
	decipheredMessage := decipher(cipheredMessage)
	if strings.Contains(decipheredMessage, substring) {
		fmt.Println(decipheredMessage)
	} else {
		fmt.Println("Invalid")
	}
}

func decipher(cipheredMessage string) string {
	decipherRatio := 2
	var decipheredWordLettersArray []rune

	for _, char := range cipheredMessage {
		if !isAlphabetical(char) {
			decipheredWordLettersArray = append(decipheredWordLettersArray, char)
		} else {
			decipheredWordLettersArray = append(decipheredWordLettersArray, decipherChar(char, decipherRatio))
		}
	}

	return string(decipheredWordLettersArray)
}

func decipherChar(char rune, decipherRatio int) rune {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	charIndex := strings.IndexRune(alphabet, toLower(char))
	decipheredIndex := (charIndex - decipherRatio + len(alphabet)) % len(alphabet)
	return rune(alphabet[decipheredIndex])
}

func isAlphabetical(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func toLower(char rune) rune {
	if char >= 'A' && char <= 'Z' {
		return char + 32
	}
	return char
}