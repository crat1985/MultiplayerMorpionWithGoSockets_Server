package main

import (
	"log"
	"net"

	"github.com/RIC217/MultiplayerMorpionWithGoSockets_Server/utils"
)

const port = "8888"

func main() {
	log.Printf("Démarrage du serveur sur le port %s...\n", port)
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Attente de connexions...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go utils.ProcessClient(conn)
	}
}
