package application

import (
	"estsoftware/src/core/security"
	"estsoftware/src/modules/domain"
	"estsoftware/src/modules/domain/entities"
	"fmt"
	"os"
	"path/filepath"
)

type CreateModule struct {
	db domain.IModule
}

func NewCreateModule(db domain.IModule) *CreateModule {
	return &CreateModule{db: db}
}

func (cm *CreateModule) Execute(module entities.Module, imageFile []byte, fileName string) (*entities.Module, error) {
	if len(imageFile) > 0 && fileName != "" {
		uploadsDir := "./image_modulo"
		fmt.Printf("Intentando guardar imagen en: %s\n", uploadsDir)

		if _, err := os.Stat(uploadsDir); os.IsNotExist(err) {
			fmt.Println("El directorio no existe, creándolo...")
			if err := os.MkdirAll(uploadsDir, 0755); err != nil {
				fmt.Printf("Error al crear el directorio: %v\n", err)
				return nil, err
			}
		}

		fileExt := filepath.Ext(fileName)
		uniqueFileName := fmt.Sprintf("%s%s", security.GenerateRandomString(8), fileExt)
		filePath := filepath.Join(uploadsDir, uniqueFileName)

		if err := os.WriteFile(filePath, imageFile, 0644); err != nil {
			fmt.Printf("Error al escribir el archivo: %v\n", err)
			return nil, err
		}
		fmt.Println("¡Imagen escrita exitosamente!")

		module.ImagenModulo = "/image_modulo/" + uniqueFileName
	}

	savedModule, err := cm.db.Save(module)
	if err != nil {
		return nil, err
	}

	return savedModule, nil
}
