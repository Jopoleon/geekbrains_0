package postgres

import (
	"github.com/Jopoleon/geekbrains_0/models"
	"github.com/pkg/errors"
)

func (db *DB) CreateNewAuthor(firstName, secondName string) error {
	_, err := db.DB.Exec("INSERT INTO authors (first_name, second_name) VALUES ($1,$2);",
		firstName, secondName)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return errors.WithStack(err)
	}
	return nil
}

func (db *DB) GetAllAuthors() ([]models.Author, error) {
	var books []models.Author
	err := db.DB.Select(&books, "SELECT * FROM authors;")
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}

	return books, nil
}
