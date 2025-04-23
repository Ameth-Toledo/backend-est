package entities

type Page struct {
	ID        int32  `json:"id"`
	Titulo    string `json:"titulo"`
	Contenido string `json:"contenido"`
	ModuloID  int32  `json:"modulo_id"`
}
