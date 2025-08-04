package store

import (
	"database/sql"
	"errors"
	"time"

	"github.com/navid-fn/shorty/internal/api/model"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int       `json:"id"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type PostgresUserStore struct {
	db *sql.DB
}

func NewPostgresUserStore(db *sql.DB) *PostgresUserStore {
	return &PostgresUserStore{db: db}
}

type UserStore interface {
	CreateUser(req *model.UserRegister) (*User, error)
	Authenticate(req *model.UserLogin) (bool, error)
}

func (pgdb *PostgresUserStore) CreateUser(userReq *model.UserRegister) (*User, error) {
	tx, err := pgdb.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	user := &User{
		UserName: userReq.Username,
		Email:    userReq.Email,
	}

	query :=
		`INSERT INTO users (username, password_hash, email)
		VALUES ($1, $2, $3)
		returning  id, created_at
	`
	err = tx.QueryRow(query, userReq.Username, hashedPassword, userReq.Email).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (pgdb *PostgresUserStore) Authenticate(req *model.UserLogin) (bool, error) {
	var hashedPassword string
	query := `
	SELECT password_hash
	FROM users 
	WHERE username = $1
	`
	err := pgdb.db.QueryRow(query, req.Username).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
	if err != nil {
		return false, errors.New("invalid credentials")
	}
	return true, nil
}
