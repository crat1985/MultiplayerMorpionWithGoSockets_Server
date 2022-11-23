package utils

import "log"

// Ajoute l'élément passé en argument au slice d'utilisateurs.
func AddToUsers(u User) {
	Users = append(Users, u)
	log.Println(u.pseudo + " ajouté aux utilisateurs !")
}

// Supprime l'utilisateur ayant le pseudo passé en argument du slice d'utilisateurs.
func RemoveFromUsersByPseudo(pseudo string) {
	//Initialise index à -1
	var index int = -1
	for i, user := range Users {
		if user.pseudo == pseudo {
			index = i
		}
	}
	//Vérifier si index a changé, donc si un utilisateur avec ce pseudo est connecté
	if index == -1 {
		log.Println("Pseudo introuvable dans les utilisateurs en ligne.")
		return
	}
	//Supprimer l'utilisateur avec ce pseudo du slice d'utilisateurs
	Users = append(Users[:index], Users[index+1:]...)
	log.Printf("%s supprimé de la liste des utilisateurs en ligne !\n", pseudo)
}

// Renvoie true si le pseudo passé en argument correspond à un utilisateur en ligne, faux sinon.
func IsConnected(pseudo string) bool {
	//Initialise connected à sa valeur par défaut (false)
	var connected bool
	//Vérifie si l'utilisateur est connecté
	for _, user := range Users {
		if user.pseudo == pseudo {
			connected = true
			break
		}
	}
	//Renvoie la value de connected donc si l'utilisateur est connecté
	return connected
}
