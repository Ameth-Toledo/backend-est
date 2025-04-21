package domain

import "estsoftware/src/users/domain/entities"

type IUser interface {
	Save(user entities.User) (entities.User, error)
	GetByCorreo(email string) (*entities.User, error)
	GetAll() ([]entities.User, error)
	GetById(id int) (entities.User, error)
	Edit(entities.User) error
	Delete(id int) error
}
