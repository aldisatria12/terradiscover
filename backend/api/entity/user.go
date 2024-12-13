package entity

import (
	"database/sql"
	"time"
)

type User struct {
	Id        int
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
