package db

import (
	"database/sql"
)

type User struct {
	ID             string `db:"id"`
	Name           string `db:"name"`
	HashedPassword string `db:"hashed_password"`
	Email          string `db:"email"`
}

func InsertUser(tx *sql.Tx, user User) (err error) {

	_, err = tx.Exec(`INSERT INTO users (id, name, hashed_password, email) VALUES ($1, $2, $3, $4)`, user.ID, user.Name, user.HashedPassword, user.Email)

	if err != nil {
		return
	}

	return
}

func GetUserByEmail(tx *sql.Tx, email string) (*User, error) {
	user := &User{}
	rows := tx.QueryRow(`SELECT * FROM users WHERE email = $1`, email)

	err := rows.Scan(&user.ID, &user.Name, &user.HashedPassword, &user.Email)
	if err != nil {
		return nil, err
	}

	return user, err
}

func GetAllUsers(tx *sql.Tx) (users []User, err error) {
	var rows *sql.Rows
	rows, err = tx.Query(`SELECT * FROM users`)
	if err != nil {
		return
	}

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.HashedPassword, &user.Email)
		if err != nil {
			return
		}

		users = append(users, user)
	}

	return
}
