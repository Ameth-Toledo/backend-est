package controllers

import (
	"crypto/rand"
	"encoding/hex"
	"estsoftware/src/materials/application"
	"estsoftware/src/materials/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type CreateMaterialController struct {
	UseCase *application.CreateMaterial
}

func NewCreateMaterialController(useCase *application.CreateMaterial) *CreateMaterialController {
	return &CreateMaterialController{UseCase: useCase}
}

func (c *CreateMaterialController) Execute(ctx *gin.Context) {
	var material entities.Material

	material.Tipo = ctx.PostForm("tipo")
	material.Enlace = ctx.PostForm("enlace")

	paginaIDStr := ctx.PostForm("pagina_id")
	paginaID, err := strconv.Atoi(paginaIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "pagina_id inválido"})
		return
	}
	material.PaginaID = paginaID

	file, err := ctx.FormFile("archivo")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Archivo no proporcionado", "detalle": err.Error()})
		return
	}

	uploadDir := "./materials"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear la carpeta de subida", "detalle": err.Error()})
		return
	}

	originalName := file.Filename
	fileExt := filepath.Ext(originalName)
	randomStr := generateRandomString(8)

	safeName := "material_" + randomStr + fileExt

	filePath := filepath.Join(uploadDir, safeName)
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el archivo", "detalle": err.Error()})
		return
	}

	material.ArchivoURL = "http://localhost:8080/static/materials/" + safeName

	created, err := c.UseCase.Execute(material)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el material", "detalle": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, created)
}

func generateRandomString(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "default"
	}
	return hex.EncodeToString(bytes)
}
