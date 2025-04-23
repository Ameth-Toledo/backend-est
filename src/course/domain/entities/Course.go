package entities

type Course struct {
	ID            int32  `json:"id"`
	Titulo        string `json:"titulo"`
	Descripcion   string `json:"descripcion"`
	ImagenPortada string `json:"imagen_portada"`
	ProfesorId    int32  `json:"profesor_id"`
	EsGratuito    bool   `json:"es_gratuito"`
}
