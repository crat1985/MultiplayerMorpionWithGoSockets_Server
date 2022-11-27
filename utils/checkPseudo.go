package utils

import (
	"strings"
	"unicode"
)

// Renvoie l'erreur en tant que string s'il y en a une, true pour le second renvoi si le pseudo est valide.
func IsPseudoValid(pseudo string) (string, bool) {
	//Vérifie si la première lettre est une lettre
	if !unicode.IsLetter(rune(pseudo[0])) {
		return "Le pseudo doit commencer par une lettre !", false
	}
	//Vérifie si le pseudo a une longueur d'au moins 5 caractères
	if len(pseudo) < 5 {
		return "Le pseudo doit contenir au moins 5 caractères !", false
	}
	//Vérifie si le pseudo entier ne contient que des lettres, des nombres et des underscores
	for _, char := range pseudo {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) && string(char) != "_" {
			return "Le pseudo doit commencer par un lettre et ne contenir que des lettres, des chiffres et des underscores !", false
		}
	}
	//Vérifie que le pseudo ne contient ni d'espaces, ni de tabulations, ni de retours à la ligne
	if strings.Contains(pseudo, "\n") || strings.Contains(pseudo, "\t") || strings.Contains(pseudo, " ") {
		return "Le pseudo ne peut pas contenir d'espaces, de tabulations ou de retours à la ligne !", false
	}
	//Renvoyer que tout s'est bien passé
	return "", true
}
