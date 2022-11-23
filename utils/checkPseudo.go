package utils

import (
	"strings"
	"unicode"
)

func IsPseudoValid(pseudo string) bool {
	if !unicode.IsLetter(rune(pseudo[0])) {
		return false
	}
	for _, char := range pseudo {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) && string(char) != "_" {
			return false
		}
	}
	if strings.Contains(pseudo, "\n") || strings.Contains(pseudo, "\t") || strings.Contains(pseudo, " ") {
		return false
	}
	return true
}
