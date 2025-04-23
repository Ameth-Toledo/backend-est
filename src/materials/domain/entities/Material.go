package entities

type Material struct {
	ID         int    `json:"id"`
	Tipo       string `json:"tipo"`
	ArchivoURL string `json:"archivo_url"`
	Enlace     string `json:"enlace"`
	PaginaID   int    `json:"pagina_id"`
}
