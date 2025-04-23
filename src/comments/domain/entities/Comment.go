package entities

import "time"

type Comment struct {
	ID        int32     `json:"id"`
	UsuarioID int32     `json:"usuario_id"`
	CursoID   *int32    `json:"curso_id"`
	Contenido string    `json:"contenido"`
	Fecha     time.Time `json:"fecha"`
}
