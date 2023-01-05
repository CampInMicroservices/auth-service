package db

import (
	"context"
	"log"
	"strconv"
	"time"
)

const (
	mySigningKey = "my-secret-key"
)

type User struct {
	ID        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	Activated bool      `json:"activated" db:"activated"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type LoginParam struct {
	Email    string
	Password string
}

type Credentials struct {
	JWT string
}

func (store *Store) AuthenticateUser(ctx context.Context, arg LoginParam) (Credentials, error) {

	const query = `SELECT * FROM "users" WHERE "email" = $1 AND "password" = $2`
	row := store.db.QueryRowContext(ctx, query, arg.Email, arg.Password)

	log.Println("sem tu")

	var user User
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Activated,
		&user.CreatedAt,
	)

	if err != nil {
		return Credentials{}, err
	}

	log.Println("sem tu 2")

	var credentias Credentials
	jwt, err := generateJWT(user.ID)

	if err != nil {
		return Credentials{}, err
	}

	credentias.JWT = jwt

	return credentias, err
}

func generateJWT(userId int64) (string, error) {
	return strconv.FormatInt(userId, 10), nil
}
