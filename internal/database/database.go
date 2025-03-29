package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/VeeRomanoff/Lollipop/internal/domain"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "username"
	password = "username666"
	dbname   = "lollipop_db"
)

type Database struct {
	db *sql.DB
}

func New() (*Database, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to obtain database connection: %s", err.Error())
	}

	return &Database{
		db: db,
	}, nil
}

func (d *Database) RegisterUser(ctx context.Context, user *domain.User) (int64, error) {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	query, args, err := sq.Insert("lollipop_users").
		Columns("name", "age", "height", "hobbies", "description").
		Values(
			user.Name,
			user.Age,
			user.Height,
			convertHobbiesIntoJSON(user.Hobbies),
			user.Description,
		).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return 0, errors.Join(domain.ErrorInternal, err)
	}

	log.Println("Executing query:", query)

	row := tx.QueryRowContext(ctx, query, args...)

	log.Printf("Executing query: %s, args: %v", query, args)

	var id int64

	if err = row.Scan(&id); err != nil {
		return 0, errors.Join(domain.ErrorInternal, err)
	}

	if err = tx.Commit(); err != nil {
		return 0, errors.Join(domain.ErrorInternal, err)
	}

	log.Printf("User created with id: %d", id)

	return id, nil
}

func convertHobbiesIntoJSON(hobbies []string) []byte {
	bytes, err := json.Marshal(hobbies)
	if err != nil {
		return []byte{}
	}

	return bytes
}
