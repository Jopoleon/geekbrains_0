package postgres

import (
	"github.com/Jopoleon/geekbrains_0/models"
	"github.com/pkg/errors"
)

func (db *DB) CreateNewBook(name string, pages int) error {
	_, err := db.DB.Exec("INSERT INTO books (name, pages) VALUES ($1,$2);",
		name, pages)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return errors.WithStack(err)
	}
	return nil
}

func (db *DB) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	err := db.DB.Select(&books, "SELECT * FROM books;")
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}

	return books, nil
}
