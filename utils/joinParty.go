package utils

import (
	"errors"
	"log"
)

func JoinParty(user User, partyId string) {
	party, err := GetPartyById(partyId)
	if err != nil {
		log.Print(err)
		return
	}
	if party.player.pseudo == "" {
		party.player = user
	} else {
		user.socket.Write([]byte("partyfull"))
	}
}

func GetPartyById(id string) (Party, error) {
	for _, party := range Parties {
		if party.id == id {
			return party, nil
		}
	}
	return Party{}, errors.New("partie introuvable")
}
