package application

import (
	"errors"
	"estsoftware/src/users/domain"
	"estsoftware/src/users/domain/entities"
)

type ViewByIdUser struct {
	db domain.IUser
}

func NewUserById(db domain.IUser) *ViewByIdUser {
	return &ViewByIdUser{db: db}
}

func (vc *ViewByIdUser) Execute(id int) (entities.User, error) {
	user, err := vc.db.GetById(id)
	if err != nil {
		return entities.User{}, errors.New("User not found")
	}
	return user, nil
}
