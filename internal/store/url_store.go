package store

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/navid-fn/shorty/internal/utils"
)

type Url struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	OriginalUrl string    `json:"original_url"`
	ShortCode   string    `json:"short_code"`
	Clicked     int       `json:"clicked"`
}

type PostgresUrlStore struct {
	db *sql.DB
}

func NewPostgresUrlStore(db *sql.DB) *PostgresUrlStore {
	return &PostgresUrlStore{db: db}
}

type UrlStore interface {
	CreateUrl(*Url) (*Url, error)
	GetOrginalUrlByString(code string) (*string, error)
	CheckDuplicateShortCode(code string) bool
}

func (pgdb *PostgresUrlStore) CreateUrl(url *Url) (*Url, error) {
	tx, err := pgdb.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query :=
		`INSERT INTO urls (original_url, short_code)
		VALUES ($1, $2)
		RETURNING id, created_at,short_code 
	`
	shortCode := utils.GeneratePseudoRandomString(5)
	for pgdb.CheckDuplicateShortCode(shortCode) {
		fmt.Println("short code", shortCode)
		shortCode = utils.GeneratePseudoRandomString(5)

	}
	err = tx.QueryRow(query, url.OriginalUrl, shortCode).Scan(&url.ID, &url.CreatedAt, &url.ShortCode)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return url, nil
}

func (pgdb *PostgresUrlStore) GetOrginalUrlByString(code string) (*string, error) {
	url := &Url{}
	query := `
	SELECT original_url
	FROM ulrs
	WHERE short_code = $1
	`
	err := pgdb.db.QueryRow(query, code).Scan(&url.OriginalUrl)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &url.OriginalUrl, nil
}

func (pgdb *PostgresUrlStore) CheckDuplicateShortCode(code string) bool {
	var exists bool
	query := `
	SELECT EXISTS
	(
		SELECT 1
		FROM ulrs
		WHERE short_code = $1
	)
	`
	err := pgdb.db.QueryRow(query, code).Scan(&exists)
	if err != nil || err == sql.ErrNoRows {
		return false
	}
	return exists
}
