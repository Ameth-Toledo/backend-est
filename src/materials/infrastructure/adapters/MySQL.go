package adapters

import (
	"database/sql"
	"estsoftware/src/materials/domain/entities"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) *MySQL {
	return &MySQL{conn: conn}
}

func (r *MySQL) Save(material entities.Material) (*entities.Material, error) {
	query := `INSERT INTO materiales (tipo, archivo_url, enlace, pagina_id) VALUES (?, ?, ?, ?)`
	result, err := r.conn.Exec(query, material.Tipo, material.ArchivoURL, material.Enlace, material.PaginaID)
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	material.ID = int(id)
	return &material, nil
}

func (r *MySQL) GetAll() ([]entities.Material, error) {
	rows, err := r.conn.Query("SELECT * FROM materiales")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var materials []entities.Material
	for rows.Next() {
		var m entities.Material
		if err := rows.Scan(&m.ID, &m.Tipo, &m.ArchivoURL, &m.Enlace, &m.PaginaID); err != nil {
			return nil, err
		}
		materials = append(materials, m)
	}
	return materials, nil
}

func (r *MySQL) GetById(id int) (*entities.Material, error) {
	row := r.conn.QueryRow("SELECT * FROM materiales WHERE id = ?", id)
	var m entities.Material
	if err := row.Scan(&m.ID, &m.Tipo, &m.ArchivoURL, &m.Enlace, &m.PaginaID); err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *MySQL) Update(material entities.Material) (*entities.Material, error) {
	query := `UPDATE materiales SET tipo=?, archivo_url=?, enlace=?, pagina_id=? WHERE id=?`
	_, err := r.conn.Exec(query, material.Tipo, material.ArchivoURL, material.Enlace, material.PaginaID, material.ID)
	if err != nil {
		return nil, err
	}
	return &material, nil
}

func (r *MySQL) Delete(id int) error {
	_, err := r.conn.Exec("DELETE FROM materiales WHERE id = ?", id)
	return err
}
