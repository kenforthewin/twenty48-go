package main

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int            `json:"id"`
	TopScore  *sql.NullInt64 `json:"topScore,omitempty"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"createdAt"`
	Password  sql.NullString `json:"-"`
}

type Users []User
