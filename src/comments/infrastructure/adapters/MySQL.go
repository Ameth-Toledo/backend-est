package adapters

import (
	"database/sql"
	"estsoftware/src/comments/domain/entities"
	"fmt"
	"time"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) *MySQL {
	return &MySQL{conn: conn}
}

func (m *MySQL) Save(comment entities.Comment) (*entities.Comment, error) {
	query := `INSERT INTO comentarios (usuario_id, curso_id, contenido) VALUES (?, ?, ?)`
	result, err := m.conn.Exec(query, comment.UsuarioID, comment.CursoID, comment.Contenido)
	if err != nil {
		return nil, fmt.Errorf("error al insertar el comentario: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error al obtener LastInsertId: %v", err)
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
		var fecha []byte // Cambiamos para leer el tipo como []byte
		err := rows.Scan(&c.ID, &c.UsuarioID, &cursoID, &c.Contenido, &fecha)
		if err != nil {
			return nil, fmt.Errorf("error al escanear los datos: %v", err)
		}

		if cursoID.Valid {
			c.CursoID = &cursoID.Int32
		}

		// Convertimos la fecha manualmente
		if len(fecha) > 0 {
			parsedFecha, err := time.Parse("2006-01-02 15:04:05", string(fecha)) // Cambiamos el formato de fecha
			if err != nil {
				return nil, fmt.Errorf("error al convertir la fecha: %v", err)
			}
			c.Fecha = parsedFecha
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
	var fechaBytes []byte

	err := row.Scan(&c.ID, &c.UsuarioID, &cursoID, &c.Contenido, &fechaBytes)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al escanear datos: %v", err)
	}

	if cursoID.Valid {
		c.CursoID = &cursoID.Int32
	}

	if len(fechaBytes) > 0 {
		parsedFecha, err := time.Parse("2006-01-02 15:04:05", string(fechaBytes))
		if err != nil {
			return nil, fmt.Errorf("error al convertir la fecha: %v", err)
		}
		c.Fecha = parsedFecha
	}

	return &c, nil
}

func (m *MySQL) Update(comment entities.Comment) (*entities.Comment, error) {
	query := `UPDATE comentarios SET usuario_id = ?, curso_id = ?, contenido = ? WHERE id = ?`
	_, err := m.conn.Exec(query, comment.UsuarioID, comment.CursoID, comment.Contenido, comment.ID)
	if err != nil {
		return nil, err
	}

	return m.GetById(int(comment.ID))
}

func (m *MySQL) Delete(id int) error {
	query := `DELETE FROM comentarios WHERE id = ?`
	_, err := m.conn.Exec(query, id)
	return err
}
