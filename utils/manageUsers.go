package utils

import "log"

// Ajoute l'élément passé en argument au slice d'utilisateurs.
func AddToUsers(u User) {
	Users = append(Users, u)
	log.Println(u.pseudo + " ajouté aux utilisateurs !")
}

// Supprime l'utilisateur ayant le pseudo passé en argument du slice d'utilisateurs.
func RemoveFromUsersByPseudo(pseudo string) {
	var index int = -1
	for i, user := range Users {
		if user.pseudo == pseudo {
			index = i
		}
	}
	if index == -1 {
		log.Println("Pseudo introuvable dans les utilisateurs en ligne.")
		return
	}
	index = -1
	Users = append(Users[:index], Users[index+1:]...)
	log.Printf("%s supprimé de la liste des utilisateurs en ligne !\n", pseudo)
}

// Renvoie true si le pseudo passé en argument correspond à un utilisateur en ligne, faux sinon.
func IsConnected(pseudo string) bool {
	var connected bool
	for _, user := range Users {
		if user.pseudo == pseudo {
			connected = true
			break
		}
	}
	return connected
}
