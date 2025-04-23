package adapters

import (
	"database/sql"
	"estsoftware/src/modules/domain/entities"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) *MySQL {
	return &MySQL{conn: conn}
}

func (m *MySQL) Save(modulo entities.Module) (*entities.Module, error) {
	query := `INSERT INTO modulos (titulo, descripcion, imagen_modulo, curso_id, profesor_id, habilitado)
	          VALUES (?, ?, ?, ?, ?, ?)`
	result, err := m.conn.Exec(query, modulo.Titulo, modulo.Descripcion, modulo.ImagenModulo, modulo.CursoId, modulo.ProfesorId, modulo.Habilitado)
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	modulo.ID = int32(id)
	return &modulo, nil
}

func (m *MySQL) GetAll() ([]entities.Module, error) {
	query := `SELECT id, titulo, descripcion, imagen_modulo, curso_id, profesor_id, habilitado FROM modulos`
	rows, err := m.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var modules []entities.Module
	for rows.Next() {
		var module entities.Module
		err := rows.Scan(
			&module.ID,
			&module.Titulo,
			&module.Descripcion,
			&module.ImagenModulo,
			&module.CursoId,
			&module.ProfesorId,
			&module.Habilitado,
		)
		if err != nil {
			return nil, err
		}
		modules = append(modules, module)
	}
	return modules, nil
}

func (m *MySQL) GetById(id int) (*entities.Module, error) {
	query := `SELECT id, titulo, descripcion, imagen_modulo, curso_id, profesor_id, habilitado FROM modulos WHERE id = ?`
	row := m.conn.QueryRow(query, id)

	var module entities.Module
	err := row.Scan(
		&module.ID,
		&module.Titulo,
		&module.Descripcion,
		&module.ImagenModulo,
		&module.CursoId,
		&module.ProfesorId,
		&module.Habilitado,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &module, nil
}

func (m *MySQL) Update(modulo entities.Module) (*entities.Module, error) {
	query := `
		UPDATE modulos
		SET titulo = ?, descripcion = ?, imagen_modulo = ?, curso_id = ?, profesor_id = ?, habilitado = ?
		WHERE id = ?
	`
	_, err := m.conn.Exec(query, modulo.Titulo, modulo.Descripcion, modulo.ImagenModulo, modulo.CursoId, modulo.ProfesorId, modulo.Habilitado, modulo.ID)
	if err != nil {
		return nil, err
	}
	return &modulo, nil
}

func (m *MySQL) Delete(id int) error {
	query := `DELETE FROM modulos WHERE id = ?`
	_, err := m.conn.Exec(query, id)
	return err
}
