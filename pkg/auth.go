package pkg

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Morakab struct {
	DB *sql.DB
}

func GeneratePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (m *Morakab) RegisterUser(username string, email string, password string) error {
	var usernameCheck string
	err := m.DB.QueryRow(`SELECT username FROM users WHERE username=$1`, username).Scan(&usernameCheck)
	if usernameCheck != "" {
		return errors.New("Username already exists")
	}
	if err != sql.ErrNoRows {
		return err
	}
	var emailCheck string
	err = m.DB.QueryRow(`SELECT email FROM users WHERE email=$1`, email).Scan(&emailCheck)
	if emailCheck != "" {
		return errors.New("Email already exists")
	}
	if err != sql.ErrNoRows {
		return err
	}
	password_hash, err := GeneratePasswordHash(password)
	if err != nil {
		return err
	}
	_, err = m.DB.Exec(`INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)`, username, email, password_hash)
	if err != nil {
		return err
	}
	return nil
}

func (m *Morakab) LoginUser(username string, password string) error {
	var password_hash string
	if err := m.DB.QueryRow(`SELECT password_hash FROM users WHERE username = $1`, username).Scan(&password_hash); err == nil {
		if !CheckPasswordHash(password, password_hash) {
			return errors.New("Invalid username or password")
		}
	} else {
		return err
	}

	return nil
}
