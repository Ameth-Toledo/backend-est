package entities

import "time"

type Inscription struct {
	ID               int       `json:"id"`
	AlumnoID         int       `json:"alumno_id"`
	CursoID          int       `json:"curso_id"`
	FechaInscripcion time.Time `json:"fecha_inscripcion"`
}
