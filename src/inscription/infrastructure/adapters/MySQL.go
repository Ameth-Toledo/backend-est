package adapters

import (
	"database/sql"
	"estsoftware/src/inscription/domain/entities"
	"fmt"
	"time"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) *MySQL {
	return &MySQL{conn: conn}
}

func (m *MySQL) Save(inscripcion entities.Inscription) (*entities.Inscription, error) {
	query := `INSERT INTO inscripciones (alumno_id, curso_id) VALUES (?, ?)`

	result, err := m.conn.Exec(query, inscripcion.AlumnoID, inscripcion.CursoID)
	if err != nil {
		return nil, fmt.Errorf("failed to insert inscripcion: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last insert ID: %v", err)
	}

	return m.GetByID(int(id))
}

func (m *MySQL) GetAll() ([]entities.Inscription, error) {
	query := `SELECT id, alumno_id, curso_id, fecha_inscripcion FROM inscripciones`

	rows, err := m.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all inscripciones: %v", err)
	}
	defer rows.Close()

	var inscripciones []entities.Inscription
	for rows.Next() {
		var i entities.Inscription
		var fechaStr []byte

		if err := rows.Scan(&i.ID, &i.AlumnoID, &i.CursoID, &fechaStr); err != nil {
			return nil, fmt.Errorf("failed to scan inscripcion: %v", err)
		}

		fecha, err := time.Parse("2006-01-02 15:04:05", string(fechaStr))
		if err != nil {
			return nil, fmt.Errorf("failed to parse fecha_inscripcion: %v", err)
		}
		i.FechaInscripcion = fecha

		inscripciones = append(inscripciones, i)
	}
	return inscripciones, nil
}

func (m *MySQL) GetByID(id int) (*entities.Inscription, error) {
	query := `SELECT id, alumno_id, curso_id, fecha_inscripcion FROM inscripciones WHERE id = ?`

	var inscripcion entities.Inscription
	var fechaStr []byte

	err := m.conn.QueryRow(query, id).Scan(
		&inscripcion.ID,
		&inscripcion.AlumnoID,
		&inscripcion.CursoID,
		&fechaStr,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get inscripcion by id: %v", err)
	}

	fecha, err := time.Parse("2006-01-02 15:04:05", string(fechaStr))
	if err != nil {
		return nil, fmt.Errorf("failed to parse fecha_inscripcion: %v", err)
	}
	inscripcion.FechaInscripcion = fecha

	return &inscripcion, nil
}

func (m *MySQL) Update(id int, inscripcion entities.Inscription) (*entities.Inscription, error) {
	query := `UPDATE inscripciones SET alumno_id = ?, curso_id = ? WHERE id = ?`

	_, err := m.conn.Exec(query, inscripcion.AlumnoID, inscripcion.CursoID, id)
	if err != nil {
		return nil, fmt.Errorf("failed to update inscripcion: %v", err)
	}

	updated, err := m.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated inscripcion: %v", err)
	}
	return updated, nil
}

func (m *MySQL) Delete(id int) error {
	query := `DELETE FROM inscripciones WHERE id = ?`
	_, err := m.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete inscripcion: %v", err)
	}
	return nil
}
