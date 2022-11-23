package utils

import (
	"errors"
	"net"
)

func AskPseudo(conn net.Conn) (string, error) {
	slice := make([]byte, 1024)
	n, err := conn.Read(slice)
	if err != nil {
		return "", err
	}
	pseudo := string(slice[:n])
	if !IsPseudoValid(pseudo) {
		conn.Write([]byte("Pseudo invalide !"))
		return "", errors.New("pseudo invalide")
	}
	return pseudo, nil
}
