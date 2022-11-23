package utils

import (
	"fmt"
	"log"
	"net"
)

func ProcessClient(conn net.Conn) {
	fmt.Println("Nouvelle connexion de " + conn.RemoteAddr().String() + " !")
	pseudo, err := AskPseudo(conn)
	if err != nil {
		log.Print(pseudo)
		return
	}
}
