// internal/models/users.go

package models

import (
	"time"
)

type User struct {
	UserID       int        `db:"user_id"`
	Username     string     `db:"username"`
	PasswordHash string     `db:"password_hash"`
	DateCreated  time.Time  `db:"date_created"`
	LastLogin    *time.Time `db:"last_login"`
	IsActive     bool       `db:"is_active"`
}
