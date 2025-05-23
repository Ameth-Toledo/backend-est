package entities

type User struct {
	ID         int32   `json:"id"`
	Nombre     string  `json:"nombre"`
	Apellido   string  `json:"apellido"`
	Correo     string  `json:"correo"`
	Contrasena string  `json:"contrasena"`
	FotoPerfil *string `json:"foto_perfil"`
	RolID      *int32  `json:"rol_id,omitempty"`
	Plan       string  `json:"plan"`
}
