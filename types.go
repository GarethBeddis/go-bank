package main

import (
	"time"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type Account struct {
	ID        uint   `gorm:"primaryKey"`
	Number    string `gorm:"type:uuid;default:gen_random_uuid()"`
	Username  string `gorm:"uniqueIndex"`
	Balance   int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func CreateAccount(username string) *Account {
	return &Account{
		Username: username,
	}
}

type CreateAccountRequest struct {
	Username string
}
