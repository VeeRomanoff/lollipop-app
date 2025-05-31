package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/VeeRomanoff/Lollipop/internal/domain"
)

const (
	usersTableName         = "lollipop_users"
	usersColumnID          = "id"
	usersColumnName        = "name"
	usersColumnAge         = "age"
	usersColumnHeight      = "height"
	usersColumnHobbies     = "hobbies"
	usersColumnDescription = "description"
	usersColumnEmail       = "email"
)

// RegisterUser регистрация пользователя
func (db *Database) RegisterUser(ctx context.Context, user *domain.User) (int64, error) {
	query, args, err := sq.Insert(usersTableName).
		Columns(usersColumnName, usersColumnAge, usersColumnHeight, usersColumnHobbies, usersColumnDescription, usersColumnEmail).
		Values(
			user.Name,
			user.Age,
			user.Height,
			marshalHobbies(user.Hobbies),
			user.Description,
			user.Email,
		).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		// TODO решить что делать с sentinel errors потому что в пакете internal/errors уже есть заготовленные sentinel errors
		return 0, fmt.Errorf("failed to build register query: %w", err)
	}

	tx, err := db.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	row := tx.QueryRowContext(ctx, query, args...)

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

func marshalHobbies(hobbies []string) []byte {
	bytes, err := json.Marshal(hobbies)
	if err != nil {
		return []byte{}
	}

	return bytes
}

// GetUserByID получение пользователя по id
func (db *Database) GetUserByID(ctx context.Context, id int64) (*domain.User, error) {
	query, args, err := queryBuilder().
		Select(
			usersColumnID,
			usersColumnName,
			usersColumnAge,
			usersColumnHeight,
			usersColumnHobbies,
			usersColumnDescription,
			usersColumnEmail,
		).
		From(usersTableName).
		Where(
			sq.Eq{"id": id},
		).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build get query: %w", err)
	}

	log.Printf("Query: %s, Args: %v", query, args) // fixme for debugging

	tx, err := db.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("beggining transaction: %w", err)
	}
	defer tx.Rollback()

	user := &domain.User{}
	var hobbiesData []byte

	if err = tx.QueryRowContext(ctx, query, args...).
		Scan(
			&user.ID,
			&user.Name,
			&user.Age,
			&user.Height,
			&hobbiesData,
			&user.Description,
			&user.Email,
		); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("user: %v", user)
			return nil, nil
		}
		log.Printf("user: %v", user)
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	user.Hobbies = unmarshalHobbies(hobbiesData)
	log.Printf("user: %v", user)
	return user, nil
}

func unmarshalHobbies(pgHobbies []byte) []string {
	var hobbies []string
	if err := json.Unmarshal(pgHobbies, &hobbies); err != nil {
		return []string{}
	}

	return hobbies
}

func (db *Database) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	query, args, err := queryBuilder().
		Update(usersTableName).
		SetMap(map[string]interface{}{
			usersColumnName:        user.Name,
			usersColumnAge:         user.Age,
			usersColumnHeight:      user.Height,
			usersColumnHobbies:     marshalHobbies(user.Hobbies),
			usersColumnDescription: user.Description,
			usersColumnEmail:       user.Email,
		}).
		Where(
			sq.Eq{
				"id": user.ID,
			},
		).Suffix("RETURNING id, name, age, height, hobbies, description, email").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build update query: %w", err)
	}

	tx, err := db.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("beggining transaction: %w", err)
	}
	defer tx.Rollback()

	row := tx.QueryRowContext(ctx, query, args...)

	var userUpdated = &domain.User{}
	var hobbiesData []byte

	if err = row.Scan(
		&userUpdated.ID,
		&userUpdated.Name,
		&userUpdated.Age,
		&userUpdated.Height,
		&hobbiesData,
		&userUpdated.Description,
		&userUpdated.Email,
	); err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	log.Println("heightWOW: ", userUpdated.Height)
	userUpdated.Hobbies = unmarshalHobbies(hobbiesData)

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("committing transaction: %w", err)
	}

	return userUpdated, nil
}

func (db *Database) DeleteUser(ctx context.Context, userID int64) error {
	tx, err := db.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("beggining transaction: %w", err)
	}
	defer tx.Rollback()

	query, args, err := queryBuilder().
		Delete(usersTableName).
		Where(
			sq.Eq{
				"id": userID,
			},
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build delete query: %w", err)
	}

	result, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute delete query: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with id %d was not found", userID)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction %w", err)
	}

	return nil
}
