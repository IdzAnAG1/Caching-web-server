package models

import (
	"time"
)

type User struct {
	ID           string
	Login        string
	PasswordHash string
	CreateAt     time.Time
}

// test
var Users = make(map[string]User)
