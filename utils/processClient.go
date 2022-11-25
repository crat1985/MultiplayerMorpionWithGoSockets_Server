package utils

import (
	"fmt"
	"log"
	"net"
	"strings"
)

// Structure contenant les informations d'un utilisateur.
type User struct {
	socket net.Conn
	pseudo string
}

type Party struct {
	owner  User
	player User
	id     string
}

var Parties []Party

// Slice de tous les utilisateurs connectés.
var Users []User

// Fonction exécutée à chaque nouvelle connexion.
func ProcessClient(conn net.Conn) {
	//Affiche l'IP du client dans la console
	fmt.Println("Nouvelle connexion de " + conn.RemoteAddr().String() + " !")
	//Appel de la fonction AskPseudo
	pseudo, err := AskPseudo(conn)
	//Gérer l'erreur s'il y en a une
	if err != nil {
		//Afficher l'erreur
		log.Print(err)
		//Fermer la connexion
		conn.Close()
		return
	}
	tempUser := User{socket: conn, pseudo: pseudo}
	AddToUsers(tempUser)
	go ListenForDatas(tempUser)
}

// Inutile et faux pour l'instant.
func ListenForDatas(user User) {
	slice := make([]byte, 1024)
	for {
		n, err := user.socket.Read(slice)
		if err != nil {
			log.Print(err)
			continue
		}
		datas := string(slice[:n])
		if datas == "createparty" {
			if !AddElementToParties(CreateParty(user)) {
				continue
			}

		}
		if strings.HasPrefix(datas, "join ") {
			JoinParty(user, strings.Split(datas, " ")[1])
			continue
		}
	}
}
