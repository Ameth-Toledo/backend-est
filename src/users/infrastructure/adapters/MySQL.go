package adapters

import (
	"database/sql"
	"errors"
	"estsoftware/src/users/domain/entities"
	"fmt"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) *MySQL {
	return &MySQL{conn: conn}
}

func (m *MySQL) Save(user entities.User) (entities.User, error) {
	query := `INSERT INTO users (nombre, apellido, correo, contrasena, foto_perfil, rol_id, plan) VALUES (?, ?, ?, ?, ?, ?, ?)`

	result, err := m.conn.Exec(
		query,
		user.Nombre,
		user.Apellido,
		user.Correo,
		user.Contrasena,
		user.FotoPerfil,
		user.RolID,
		user.Plan,
	)
	if err != nil {
		return entities.User{}, fmt.Errorf("failed to save user: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entities.User{}, fmt.Errorf("failed to retrieve last insert ID: %v", err)
	}

	user.ID = int32(id)
	return user, nil
}

func (m *MySQL) GetByEmail(correo string) (entities.User, error) {
	var user entities.User
	query := `SELECT id, nombre, apellido, correo, contrasena, foto_perfil, rol_id, plan FROM users WHERE correo = ? LIMIT 1`

	err := m.conn.QueryRow(query, correo).Scan(
		&user.ID, &user.Nombre, &user.Apellido, &user.Correo, &user.Contrasena, &user.FotoPerfil, &user.RolID, &user.Plan,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.User{}, errors.New("user not found")
		}
		return entities.User{}, err
	}

	return user, nil
}

func (m *MySQL) GetAll() ([]entities.User, error) {
	query := "SELECT id, nombre, apellido, correo, contrasena, foto_perfil, rol_id, plan FROM users"
	rows, err := m.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve users: %v", err)
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var user entities.User
		err := rows.Scan(&user.ID, &user.Nombre, &user.Apellido, &user.Correo, &user.Contrasena, &user.FotoPerfil, &user.RolID, &user.Plan)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return users, nil
}

func (m *MySQL) GetById(id int) (entities.User, error) {
	query := "SELECT id, nombre, apellido, correo, contrasena, foto_perfil, rol_id, plan FROM users WHERE id = ?"
	row := m.conn.QueryRow(query, id)

	var user entities.User
	err := row.Scan(&user.ID, &user.Nombre, &user.Apellido, &user.Correo, &user.Contrasena, &user.FotoPerfil, &user.RolID, &user.Plan)
	if err == sql.ErrNoRows {
		return entities.User{}, errors.New("user not found")
	} else if err != nil {
		return entities.User{}, fmt.Errorf("failed to retrieve user: %v", err)
	}

	return user, nil
}

func (m *MySQL) Edit(user entities.User) error {
	query := `UPDATE users SET nombre = ?, apellido = ?, correo = ?, contrasena = ?, foto_perfil = ?, rol_id = ?, plan = ? WHERE id = ?`
	_, err := m.conn.Exec(query, user.Nombre, user.Apellido, user.Correo, user.Contrasena, user.FotoPerfil, user.RolID, user.Plan, user.ID)

	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}
	return nil
}

func (m *MySQL) Delete(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := m.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}
	return nil
}

func (m *MySQL) GetByCorreo(correo string) (*entities.User, error) {
	var user entities.User
	query := `SELECT id, nombre, apellido, correo, contrasena, foto_perfil, rol_id, plan FROM users WHERE correo = ? LIMIT 1`

	err := m.conn.QueryRow(query, correo).Scan(
		&user.ID, &user.Nombre, &user.Apellido, &user.Correo, &user.Contrasena, &user.FotoPerfil, &user.RolID, &user.Plan,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("error al obtener el usuario por correo: %v", err)
	}

	return &user, nil
}
