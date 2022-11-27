package utils

import (
	"log"

	"github.com/google/uuid"
)

func CreateParty(user User) (okay bool) {
	id := uuid.New().String()
	if HasPseudoAlreadyAParty(user.pseudo) {
		user.socket.Write([]byte(""))
		return false
	}
	tempParty := Party{owner: user, id: id}
	Parties = append(Parties, tempParty)
	log.Println("Partie " + tempParty.id + " créée (owner : " + tempParty.owner.pseudo + ") !")
	return true
}
