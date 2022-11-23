package utils

import (
	"fmt"
	"log"
	"net"
)

type User struct {
	socket net.Conn
	pseudo string
}

var Users []User

func ProcessClient(conn net.Conn) {
	fmt.Println("Nouvelle connexion de " + conn.RemoteAddr().String() + " !")
	pseudo, err := AskPseudo(conn)
	if err != nil {
		log.Print(err)
		return
	}
	AddToUsers(User{socket: conn, pseudo: pseudo})
}

func ListenForDatas(conn net.Conn) {
	slice := make([]byte, 1024)
	for {
		n, err := conn.Read(slice)
		if err != nil {
			log.Print(err)
			continue
		}
		datas := string(slice[:n])
		log.Println(datas)
	}
}
