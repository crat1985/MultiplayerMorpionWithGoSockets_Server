package utils

import (
	"log"

	"github.com/google/uuid"
)

func CreateParty(user User) {
	id := uuid.New().String()
	tempParty := Party{owner: user, id: id}
	Parties = append(Parties, tempParty)
	log.Println("Partie " + tempParty.id + " créée (owner : " + tempParty.owner.pseudo + ") !")
	user.socket.Write([]byte("partycreated"))
}
