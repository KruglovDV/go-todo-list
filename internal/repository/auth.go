package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"rest-api-postgres/internal/models"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) Authorization {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user models.User) (int, error) {
	var id int

	query := fmt.Sprintf("insert into %s (name, username, password_hash) values ($1, $2, $3) returning id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthRepository) GetUser(username, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("select id from %s where username=$1 and password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}
