package controllers

import (
	"estsoftware/src/modules/application"
	"estsoftware/src/modules/domain/entities"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

type CreateModuleController struct {
	useCase *application.CreateModule
}

func NewCreateModuleController(useCase *application.CreateModule) *CreateModuleController {
	return &CreateModuleController{useCase: useCase}
}

func (cmc *CreateModuleController) Execute(c *gin.Context) {
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo procesar el formulario"})
		return
	}

	var module entities.Module
	module.Titulo = strings.TrimSpace(c.PostForm("titulo"))
	module.Descripcion = strings.TrimSpace(c.PostForm("descripcion"))

	cursoIDStr := strings.TrimSpace(c.PostForm("curso_id"))
	cursoID, err := strconv.Atoi(cursoIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "curso_id inválido"})
		return
	}
	module.CursoId = int32(cursoID)

	profesorIDStr := strings.TrimSpace(c.PostForm("profesor_id"))
	profesorID, err := strconv.Atoi(profesorIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "profesor_id inválido"})
		return
	}
	module.ProfesorId = int32(profesorID)

	habilitadoStr := strings.TrimSpace(c.PostForm("habilitado"))
	module.Habilitado = habilitadoStr == "true"

	var imageFile []byte
	var fileName string

	file, header, err := c.Request.FormFile("imagen_modulo")
	if err == nil && file != nil {
		defer file.Close()

		fileExt := filepath.Ext(header.Filename)
		if !isValidImageExt(fileExt) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de imagen no válido. Use JPG, JPEG, PNG o SVG"})
			return
		}

		imageFile, err = io.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer la imagen"})
			return
		}
		fileName = header.Filename
	}

	createdModule, err := cmc.useCase.Execute(module, imageFile, fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"modulo": createdModule})
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
