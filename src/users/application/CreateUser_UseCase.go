package application

import (
	"estsoftware/src/core/security"
	"estsoftware/src/users/domain"
	"estsoftware/src/users/domain/entities"
	"fmt"
	"os"
	"path/filepath"
)

type CreateUser struct {
	db domain.IUser
}

func NewCreateUser(db domain.IUser) *CreateUser {
	return &CreateUser{db: db}
}

func (cc *CreateUser) Execute(client entities.User, photoFile []byte, fileName string) (entities.User, error) {
	existingUser, err := cc.db.GetByCorreo(client.Correo)
	if err != nil {
		return entities.User{}, err
	}
	if existingUser != nil {
		return entities.User{}, fmt.Errorf("el correo ya está en uso, por favor use otro")
	}

	if client.Plan != "gratuito" && client.Plan != "premium" {
		return entities.User{}, fmt.Errorf("el valor del plan es inválido, debe ser 'gratuito' o 'premium'")
	}

	hashedPassword, err := security.HashPassword(client.Contrasena)
	if err != nil {
		return entities.User{}, err
	}
	client.Contrasena = hashedPassword

	if len(photoFile) > 0 && fileName != "" {
		uploadsDir := "./uploads"
		if _, err := os.Stat(uploadsDir); os.IsNotExist(err) {
			if err := os.MkdirAll(uploadsDir, 0755); err != nil {
				return entities.User{}, err
			}
		}

		fileExt := filepath.Ext(fileName)
		uniqueFileName := fmt.Sprintf("%s_%s%s", client.Correo, security.GenerateRandomString(8), fileExt)
		filePath := filepath.Join(uploadsDir, uniqueFileName)

		if err := os.WriteFile(filePath, photoFile, 0644); err != nil {
			return entities.User{}, err
		}

		photoPath := "/uploads/" + uniqueFileName
		client.FotoPerfil = &photoPath
	} else {
		defaultPath := "/img/default-avatar.png"
		client.FotoPerfil = &defaultPath
	}

	savedUser, err := cc.db.Save(client)
	if err != nil {
		return entities.User{}, err
	}

	return savedUser, nil
}
