// internal/storage/users/users.go

package users

import (
	"errors"
	"database/sql"
	"github.com/DaveVED/onion-gate/internal/models"
)

func InsertUser(db *sql.DB, user *models.User) error {
	query := `INSERT INTO users(username, password_hash, date_created, last_login, is_active) VALUES(?, ?, ?, ?, ?)`
	_, err := db.Exec(query, user.Username, user.PasswordHash, user.DateCreated, user.LastLogin, user.IsActive)
	return err
}

func FetchUserPassword(db *sql.DB, username string) (string, error) {
	var passwordHash string
	err := db.QueryRow("SELECT password_hash FROM users WHERE username = ?", username).Scan(&passwordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("no user found with the provided username")
		}
		return "", err
	}
	return passwordHash, nil
}