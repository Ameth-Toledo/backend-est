package entities

type Module struct {
	ID           int32  `json:"id"`
	Titulo       string `json:"titulo"`
	Descripcion  string `json:"descripcion"`
	ImagenModulo string `json:"imagen_modulo"`
	CursoId      int32  `json:"curso_id"`
	ProfesorId   int32  `json:"profesor_id"`
	Habilitado   bool   `json:"habilitado"`
}
