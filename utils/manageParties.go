package utils

func RemoveElementsFromPartiesByOwnerPseudo(ownerPseudo string) {
	for i, party := range Parties {
		if party.owner.pseudo == ownerPseudo {
			Parties = append(Parties[:i], Parties[i+1:]...)
		}
	}
}
