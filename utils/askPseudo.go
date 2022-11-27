package utils

import (
	"errors"
	"net"
)

// Attend le pseudo de l'utilisateur, le vérifie et revoie nil en tant qu'erreur si le pseudo est valide
func AskPseudo(conn net.Conn) (string, error) {
	//Création d'une variable qui servira à stocker les données reçues
	slice := make([]byte, 1024)
	//Attente et lecture des données envoyées par le client
	n, err := conn.Read(slice)
	//Gérer l'erreur
	if err != nil {
		return "", err
	}
	//Convertir la variable en string et la stocker dans une variable pseudo
	pseudo := string(slice[:n])
	//Vérifier si le pseudo est valide en appelant la fonction IsPseudoValid
	response, valid := IsPseudoValid(pseudo)
	//Gérer le résultat de la fonction
	if !valid {
		//Envoyer la réponse renvoyée par la fonction au client
		conn.Write([]byte(response))
		return "", errors.New("pseudo invalide")
	}
	conn.Write([]byte("pseudook"))
	return pseudo, nil
}
