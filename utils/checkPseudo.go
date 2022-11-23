package utils

import (
	"strings"
	"unicode"
)

// Renvoie l'erreur en tant que string s'il y en a une, true pour le second renvoi si le pseudo est valide.
func IsPseudoValid(pseudo string) (string, bool) {
	if !unicode.IsLetter(rune(pseudo[0])) {
		return "Le pseudo doit commencer par une lettre !", false
	}
	for _, char := range pseudo {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) && string(char) != "_" {
			return "Le pseudo doit commencer par un lettre et ne contenir que des lettres, des chiffres et des underscores !", false
		}
	}
	if strings.Contains(pseudo, "\n") || strings.Contains(pseudo, "\t") || strings.Contains(pseudo, " ") {
		return "Le pseudo ne peut pas contenir d'espaces, de tabulations ou de retours à la ligne !", false
	}
	if IsConnected(pseudo) {
		return "Ce pseudo est déjà utilisé par quelqu'un en ligne !", false
	}
	return "", true
}
