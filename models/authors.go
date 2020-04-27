package models

import "time"

type Author struct {
	FirstName  string    `json:"first_name" db:"first_name"`
	SecondName string    `json:"second_name" db:"second_name"`
	CreatedOn  time.Time `json:"created_on" db:"created_on"`
}
