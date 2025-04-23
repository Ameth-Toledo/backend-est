package controllers

import (
	"estsoftware/src/course/application"
	"estsoftware/src/course/domain/entities"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

type CreateCourseController struct {
	useCase *application.CreateCourse
}

func NewCreateCourseController(useCase *application.CreateCourse) *CreateCourseController {
	return &CreateCourseController{useCase: useCase}
}

func (cc_c *CreateCourseController) Execute(c *gin.Context) {
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo procesar el formulario"})
		return
	}

	var course entities.Course
	course.Titulo = strings.TrimSpace(c.PostForm("titulo"))
	course.Descripcion = strings.TrimSpace(c.PostForm("descripcion"))
	course.EsGratuito = strings.TrimSpace(c.PostForm("es_gratuito")) == "true"

	profesorIDStr := strings.TrimSpace(c.PostForm("profesor_id"))
	profesorID, err := strconv.Atoi(profesorIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "profesor_id inválido"})
		return
	}
	course.ProfesorId = int32(profesorID)

	var imageFile []byte
	var fileName string

	file, header, err := c.Request.FormFile("imagen_portada")
	if err == nil && file != nil {
		defer file.Close()

		fileExt := filepath.Ext(header.Filename)
		if !isValidImageExt(fileExt) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de imagen no válido. Use JPG, JPEG, PNG o SVG"})
			return
		}

		imageFile, err = io.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer el archivo"})
			return
		}
		fileName = header.Filename
	}

	createdCourse, err := cc_c.useCase.Execute(course, imageFile, fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"curso": createdCourse})
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
