package application

import (
	"errors"
	"estsoftware/src/core/security"
	"estsoftware/src/users/domain"
	"estsoftware/src/users/domain/entities"
	"fmt"
	"strings"
)

type AuthService struct {
	clientRepo domain.IUser
}

func NewAuthService(clientRepo domain.IUser) *AuthService {
	return &AuthService{clientRepo: clientRepo}
}

func (as *AuthService) Login(email, password string) (map[string]interface{}, error) {
	email = strings.TrimSpace(email)
	fmt.Println("🔍 Buscando usuario con correo:", email)

	client, err := as.clientRepo.GetByCorreo(email)
	if err != nil {
		fmt.Println("❌ Error al obtener usuario:", err)
	}
	if client == nil {
		fmt.Println("⚠ Usuario no encontrado (nil)")
		return nil, errors.New("usuario no encontrado")
	}

	if !security.CheckPassword(client.Contrasena, password) {
		fmt.Println("❌ Contraseña incorrecta")
		return nil, errors.New("contraseña incorrecta")
	}

	token, err := security.GenerateJWT(int(client.ID), client.Correo)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token":  token,
		"userId": client.ID,
		"name":   client.Nombre,
		"email":  client.Correo,
	}, nil
}

func (as *AuthService) Register(client entities.User) (entities.User, error) {
	hashedPassword, err := security.HashPassword(client.Contrasena)
	if err != nil {
		return entities.User{}, err
	}
	client.Contrasena = hashedPassword

	savedUser, err := as.clientRepo.Save(client)
	if err != nil {
		return entities.User{}, err
	}

	return savedUser, nil
}
