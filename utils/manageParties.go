package utils

func HasPseudoAlreadyAParty(pseudo string) bool {
	for _, p := range Parties {
		if p.owner.pseudo == pseudo || p.player.pseudo == pseudo {
			return true
		}
	}
	return false
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
