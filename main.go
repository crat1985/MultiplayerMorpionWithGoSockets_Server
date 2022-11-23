package main

import (
	"log"
	"net"

	"github.com/RIC217/MultiplayerMorpionWithGoSockets_Server/utils"
)

// Port d'écoute du serveur
const port = "8888"

// Main function
func main() {
	//Démarre le serveur
	log.Printf("Démarrage du serveur sur le port %s...\n", port)
	listener, err := net.Listen("tcp", ":"+port)
	//Traite l'erreur s'il y en a une
	if err != nil {
		log.Panic(err)
	}
	log.Println("Attente de connexions...")
	for {
		//Attend une connection
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		//Passe la connexion récupérée en argument de la fonction suivante (appelée en tant que Goroutine)
		go utils.ProcessClient(conn)
		//Recommence indéfiniment
	}
}
