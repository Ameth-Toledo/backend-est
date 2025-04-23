package adapters

import (
	"database/sql"
	"estsoftware/src/comments/domain/entities"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) *MySQL {
	return &MySQL{conn: conn}
}

func (m *MySQL) Save(comment entities.Comment) (*entities.Comment, error) {
	query := `INSERT INTO comentarios (usuario_id, curso_id, contenido, fecha) VALUES (?, ?, ?, ?)`
	result, err := m.conn.Exec(query, comment.UsuarioID, comment.CursoID, comment.Contenido, comment.Fecha)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	comment.ID = int32(id)
	return &comment, nil
}

func (m *MySQL) GetAll() ([]entities.Comment, error) {
	query := `SELECT id, usuario_id, curso_id, contenido, fecha FROM comentarios`
	rows, err := m.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []entities.Comment
	for rows.Next() {
		var c entities.Comment
		var cursoID sql.NullInt32
		err := rows.Scan(&c.ID, &c.UsuarioID, &cursoID, &c.Contenido, &c.Fecha)
		if err != nil {
			return nil, err
		}
		if cursoID.Valid {
			c.CursoID = &cursoID.Int32
		}
		comments = append(comments, c)
	}
	return comments, nil
}

func (m *MySQL) GetById(id int) (*entities.Comment, error) {
	query := `SELECT id, usuario_id, curso_id, contenido, fecha FROM comentarios WHERE id = ?`
	row := m.conn.QueryRow(query, id)

	var c entities.Comment
	var cursoID sql.NullInt32
	err := row.Scan(&c.ID, &c.UsuarioID, &cursoID, &c.Contenido, &c.Fecha)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	if cursoID.Valid {
		c.CursoID = &cursoID.Int32
	}
	return &c, nil
}

func (m *MySQL) Update(comment entities.Comment) (*entities.Comment, error) {
	query := `UPDATE comentarios SET usuario_id = ?, curso_id = ?, contenido = ? WHERE id = ?`
	_, err := m.conn.Exec(query, comment.UsuarioID, comment.CursoID, comment.Contenido, comment.ID)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (m *MySQL) Delete(id int) error {
	query := `DELETE FROM comentarios WHERE id = ?`
	_, err := m.conn.Exec(query, id)
	return err
}
