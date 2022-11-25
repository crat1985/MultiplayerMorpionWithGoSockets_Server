package utils

import "log"

func HasOwnerAlreadyAParty(party Party) bool {
	for _, p := range Parties {
		if p.owner.pseudo == party.owner.pseudo || p.player.pseudo == party.owner.pseudo {
			return true
		}
	}
	return false
}

func AddElementToParties(party Party) (okay bool) {
	if HasOwnerAlreadyAParty(party) {
		return false
	}
	Parties = append(Parties, party)
	log.Println("Partie " + party.id + " créée !")
	return true
}

func RemoveElementFromPartiesByOwner(owner User) {
	index := -1
	for i, party := range Parties {
		if party.owner == owner {
			index = i
			break
		}
	}
	RemoveElementFromPartiesByIndex(index)
}

func RemoveElementFromPartiesByOwnerPseudo(ownerPseudo string) {
	index := -1
	for i, party := range Parties {
		if party.owner.pseudo == ownerPseudo {
			index = i
			break
		}
	}
	RemoveElementFromPartiesByIndex(index)
}

func RemoveElementFromPartiesByParty(party Party) {
	index := -1
	for i, p := range Parties {
		if p == party {
			index = i
			break
		}
	}
	RemoveElementFromPartiesByIndex(index)
}

func RemoveElementFromPartiesByIndex(index int) {
	if index == -1 {
		return
	}
	Parties = append(Parties[:index], Parties[index+1:]...)
}
