package utils

func RemoveElementFromPartiesByOwnerPseudo(ownerPseudo string) {
	index := -1
	for i, party := range Parties {
		if party.owner.pseudo == ownerPseudo {
			index = i
			break
		}
	}
	if index == -1 {
		return
	}
	Parties = append(Parties[:index], Parties[index+1:]...)
}
