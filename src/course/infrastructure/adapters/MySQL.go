package adapters

import (
	"database/sql"
	"estsoftware/src/course/domain/entities"
	"fmt"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) *MySQL {
	return &MySQL{conn: conn}
}

func (m *MySQL) Save(course entities.Course) (*entities.Course, error) {
	query := `INSERT INTO cursos (titulo, descripcion, imagen_portada, profesor_id, es_gratuito) VALUES (?, ?, ?, ?, ?)`

	result, err := m.conn.Exec(
		query,
		course.Titulo,
		course.Descripcion,
		course.ImagenPortada,
		course.ProfesorId,
		course.EsGratuito,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to save course: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve last insert ID: %v", err)
	}

	course.ID = int32(id)
	return &course, nil
}

func (m *MySQL) GetAll() ([]entities.Course, error) {
	query := `SELECT id, titulo, descripcion, imagen_portada, profesor_id, es_gratuito FROM cursos`

	rows, err := m.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all courses: %v", err)
	}
	defer rows.Close()

	var courses []entities.Course
	for rows.Next() {
		var course entities.Course
		if err := rows.Scan(
			&course.ID,
			&course.Titulo,
			&course.Descripcion,
			&course.ImagenPortada,
			&course.ProfesorId,
			&course.EsGratuito,
		); err != nil {
			return nil, fmt.Errorf("failed to scan course: %v", err)
		}
		courses = append(courses, course)
	}
	return courses, nil
}

func (m *MySQL) GetById(id int) (*entities.Course, error) {
	query := `SELECT id, titulo, descripcion, imagen_portada, profesor_id, es_gratuito FROM cursos WHERE id = ?`

	var course entities.Course
	err := m.conn.QueryRow(query, id).Scan(
		&course.ID,
		&course.Titulo,
		&course.Descripcion,
		&course.ImagenPortada,
		&course.ProfesorId,
		&course.EsGratuito,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get course by id: %v", err)
	}
	return &course, nil
}

func (m *MySQL) GetByName(name string) ([]entities.Course, error) {
	query := `SELECT id, titulo, descripcion, imagen_portada, profesor_id, es_gratuito FROM cursos WHERE titulo LIKE ? LIMIT 10`

	rows, err := m.conn.Query(query, "%"+name+"%")
	if err != nil {
		return nil, fmt.Errorf("failed to search courses by name: %v", err)
	}
	defer rows.Close()

	var courses []entities.Course
	for rows.Next() {
		var course entities.Course
		if err := rows.Scan(
			&course.ID,
			&course.Titulo,
			&course.Descripcion,
			&course.ImagenPortada,
			&course.ProfesorId,
			&course.EsGratuito,
		); err != nil {
			return nil, fmt.Errorf("failed to scan course: %v", err)
		}
		courses = append(courses, course)
	}
	return courses, nil
}

func (m *MySQL) Update(course entities.Course) error {
	query := `UPDATE cursos SET titulo = ?, descripcion = ?, imagen_portada = ?, profesor_id = ?, es_gratuito = ? WHERE id = ?`

	_, err := m.conn.Exec(
		query,
		course.Titulo,
		course.Descripcion,
		course.ImagenPortada,
		course.ProfesorId,
		course.EsGratuito,
		course.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update course: %v", err)
	}
	return nil
}

func (m *MySQL) Delete(id int) error {
	query := `DELETE FROM cursos WHERE id = ?`

	_, err := m.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete course: %v", err)
	}
	return nil
}
