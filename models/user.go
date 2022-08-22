package models

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uuid.UUID
	Email        string
	PasswordHash string
}

type UserService struct {
	DB *sql.DB
}

type UserCreateInput struct {
	Email    string
	Password string
}

func (us *UserService) Create(input UserCreateInput) (*User, error) {
	var user User
	user.ID = uuid.Must(uuid.NewRandom())

	user.Email = strings.ToLower(input.Email)

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	user.PasswordHash = string(hashedBytes)

	row := us.DB.QueryRow(`
		INSERT INTO users (id, email, password_hash)
		VALUES ($1, $2, $3) RETURNING id`, user.ID, user.Email, user.PasswordHash,
	)
	err = row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return &user, nil
}

type UserAuthenticateInput struct {
	Email    string
	Password string
}

func (us *UserService) Authenticate(input UserAuthenticateInput) (*User, error) {
	email := strings.ToLower(input.Email)
	user := User{
		Email: email,
	}
	row := us.DB.QueryRow(`
	SELECT id, password_hash
	FROM users WHERE email=$1`, email)
	err := row.Scan(&user.ID, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("authenticate: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password))
	if err != nil {
		return nil, fmt.Errorf("authenticate: %w", err)
	}

	return &user, nil
}
