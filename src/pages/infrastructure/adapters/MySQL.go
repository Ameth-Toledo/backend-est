package adapters

import (
	"database/sql"
	"estsoftware/src/pages/domain/entities"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) *MySQL {
	return &MySQL{conn: conn}
}

func (m *MySQL) Save(page entities.Page) (*entities.Page, error) {
	query := `INSERT INTO paginas (titulo, contenido, modulo_id) VALUES (?, ?, ?)`
	result, err := m.conn.Exec(query, page.Titulo, page.Contenido, page.ModuloID)
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	page.ID = int32(id)
	return &page, nil
}

func (m *MySQL) GetAll() ([]entities.Page, error) {
	rows, err := m.conn.Query(`SELECT id, titulo, contenido, modulo_id FROM paginas`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pages []entities.Page
	for rows.Next() {
		var p entities.Page
		err := rows.Scan(&p.ID, &p.Titulo, &p.Contenido, &p.ModuloID)
		if err != nil {
			return nil, err
		}
		pages = append(pages, p)
	}
	return pages, nil
}

func (m *MySQL) GetById(id int) (*entities.Page, error) {
	row := m.conn.QueryRow(`SELECT id, titulo, contenido, modulo_id FROM paginas WHERE id = ?`, id)
	var p entities.Page
	err := row.Scan(&p.ID, &p.Titulo, &p.Contenido, &p.ModuloID)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (m *MySQL) Update(page entities.Page) (*entities.Page, error) {
	query := `UPDATE paginas SET titulo = ?, contenido = ?, modulo_id = ? WHERE id = ?`
	_, err := m.conn.Exec(query, page.Titulo, page.Contenido, page.ModuloID, page.ID)
	if err != nil {
		return nil, err
	}
	return &page, nil
}

func (m *MySQL) Delete(id int) error {
	_, err := m.conn.Exec(`DELETE FROM paginas WHERE id = ?`, id)
	return err
}
