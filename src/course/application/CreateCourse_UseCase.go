package application

import (
	"estsoftware/src/core/security"
	"estsoftware/src/course/domain"
	"estsoftware/src/course/domain/entities"
	"fmt"
	"os"
	"path/filepath"
)

type CreateCourse struct {
	db domain.ICourse
}

func NewCreateCourse(db domain.ICourse) *CreateCourse {
	return &CreateCourse{db: db}
}

func (cc *CreateCourse) Execute(course entities.Course, imageFile []byte, fileName string) (*entities.Course, error) {
	if len(imageFile) > 0 && fileName != "" {
		uploadsDir := "./image_portada"
		fmt.Printf("Intentando guardar archivo en: %s\n", uploadsDir)

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
		fmt.Println("¡Archivo escrito exitosamente!")

		imagePath := "/image_portada/" + uniqueFileName
		course.ImagenPortada = imagePath
	}

	savedCourse, err := cc.db.Save(course)
	if err != nil {
		return nil, err
	}

	return savedCourse, nil
}
