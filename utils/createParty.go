package utils

import (
	"github.com/google/uuid"
)

func CreateParty(user User) Party {
	id := uuid.New().String()
	return Party{owner: user, id: id}
}
