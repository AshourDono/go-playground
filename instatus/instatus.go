package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	decipherAndPrint("Eqfg Au hwp", "Code")
	decipherAndPrint("eqfg# Ku hwp!", "code")
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
		if !unicode.IsLetter(char) {
			decipheredWordLettersArray = append(decipheredWordLettersArray, char)
		} else {
			decipheredChar := decipherChar(char, decipherRatio)
			if unicode.IsUpper(char) {
				decipheredChar = unicode.ToUpper(decipheredChar)
			}
			decipheredWordLettersArray = append(decipheredWordLettersArray, decipheredChar)
		}
	}

	return string(decipheredWordLettersArray)
}

func decipherChar(char rune, decipherRatio int) rune {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	charIndex := strings.IndexRune(alphabet, unicode.ToLower(char))
	decipheredIndex := (charIndex - decipherRatio + len(alphabet)) % len(alphabet)
	return rune(alphabet[decipheredIndex])
}
