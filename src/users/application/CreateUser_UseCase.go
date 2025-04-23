package application

import (
	"estsoftware/src/core/security"
	"estsoftware/src/users/domain"
	"estsoftware/src/users/domain/entities"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
		fmt.Printf("Intentando guardar archivo en: %s\n", uploadsDir)

		absPath, err := filepath.Abs(uploadsDir)
		if err != nil {
			fmt.Printf("Error al obtener la ruta absoluta: %v\n", err)
		} else {
			fmt.Printf("Ruta absoluta: %s\n", absPath)
		}

		if _, err := os.Stat(uploadsDir); os.IsNotExist(err) {
			fmt.Println("El directorio no existe, creándolo...")
			if err := os.MkdirAll(uploadsDir, 0755); err != nil {
				fmt.Printf("Error al crear el directorio: %v\n", err)
				return entities.User{}, err
			}
		}

		safeEmail := strings.ReplaceAll(client.Correo, "@", "_at_")
		safeEmail = strings.ReplaceAll(safeEmail, ".", "_dot_")

		fileExt := filepath.Ext(fileName)
		uniqueFileName := fmt.Sprintf("%s_%s%s", safeEmail, security.GenerateRandomString(8), fileExt)
		filePath := filepath.Join(uploadsDir, uniqueFileName)

		fmt.Printf("Escribiendo archivo en: %s\n", filePath)
		fmt.Printf("Tamaño del archivo: %d bytes\n", len(photoFile))

		if err := os.WriteFile(filePath, photoFile, 0644); err != nil {
			fmt.Printf("Error al escribir el archivo: %v\n", err)
			return entities.User{}, err
		}
		fmt.Println("¡Archivo escrito exitosamente!")

		photoPath := "/uploads/" + uniqueFileName
		client.FotoPerfil = &photoPath
	} else {
		defaultPath := "/img/default-avatar.png"
		client.FotoPerfil = &defaultPath
	}

	fmt.Printf("Valor de FotoPerfil antes de guardar: %v\n", client.FotoPerfil)

	savedUser, err := cc.db.Save(client)
	if err != nil {
		return entities.User{}, err
	}

	fmt.Printf("Valor de FotoPerfil después de guardar: %v\n", savedUser.FotoPerfil)

	return savedUser, nil
}
