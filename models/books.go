package models

import "time"

type Book struct {
	Name      string    `json:"name" db:"name"`
	Pages     int       `json:"pages" db:"pages"`
	CreatedOn time.Time `json:"created_on" db:"created_on"`
}
