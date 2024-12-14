package entity

import (
	"database/sql"
	"time"
)

type Contact struct {
	Id        int
	UserId    int
	Name      string
	Phone     string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
