package store

import (
	"database/sql"
	"time"
)


type Url struct {
	ID int
	CreatedAt time.Time
	OrginalUrl string
	ShortCode string
	Clicked int
}

type PostgresUrlStore struct {
	db *sql.DB
}

func NewPostgresPostStore(db *sql.DB) *PostgresUrlStore {
	return &PostgresUrlStore{db: db}
}

type UrlStore interface {
	CreatetUrl(*Url) (*Url, error)
	GetUrlByID(id int64) (*Url, error)
	CheckDuplicateShortCode(code string)(bool, error)
}

func (pgdb *PostgresUrlStore)CreatetUrl(url *Url) (*Url, error) {
	tx, err := pgdb.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query :=
		`INSERT INTO urls (orginal_url, short_code)
		VALUES ($1, $2)
		RETURNING id, created_at
	`
	err = tx.QueryRow(query, url.OrginalUrl).Scan(&url.ID, &url.CreatedAt)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return url, nil
}


func (pgdb *PostgresUrlStore)CheckDuplicateShortCode() bool{
	return true
}
