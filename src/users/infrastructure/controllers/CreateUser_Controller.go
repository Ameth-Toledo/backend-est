package controllers

import (
	"estsoftware/src/users/application"
	"estsoftware/src/users/domain/entities"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"path/filepath"
	"strings"
)

type CreateUserController struct {
	useCase application.CreateUser
}

func NewCreateUserController(useCase application.CreateUser) *CreateUserController {
	return &CreateUserController{useCase: useCase}
}

func (cc_c *CreateUserController) Execute(c *gin.Context) {
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo procesar el formulario"})
		return
	}

	var input entities.User
	input.Nombre = strings.TrimSpace(c.PostForm("nombre"))
	input.Correo = strings.TrimSpace(c.PostForm("correo"))
	input.Contrasena = strings.TrimSpace(c.PostForm("contrasena"))
	input.Plan = strings.TrimSpace(c.PostForm("plan"))
	if input.Plan == "" {
		input.Plan = "gratuito"
	}

	var photoFile []byte
	var fileName string

	file, header, err := c.Request.FormFile("foto_perfil")
	if err == nil {
		defer file.Close()

		fileExt := filepath.Ext(header.Filename)
		if !isValidImageExt(fileExt) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de imagen no válido. Use JPG, JPEG, PNG o SVG"})
			return
		}

		photoFile, err = io.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer el archivo"})
			return
		}
		fileName = header.Filename
	}

	if input.Nombre == "" || input.Correo == "" || input.Contrasena == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Falta información requerida"})
		return
	}

	savedUser, err := cc_c.useCase.Execute(input, photoFile, fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	savedUser.Contrasena = ""

	c.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

func isValidImageExt(ext string) bool {
	validExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".svg":  true,
	}
	return validExts[strings.ToLower(ext)]
}
