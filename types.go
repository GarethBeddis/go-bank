package main

import (
	"time"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type Account struct {
	ID        uint  `gorm:"primaryKey"`
	Number    int64 `gorm:"type:uuid;default:gen_random_uuid()"`
	Username  string
	Balance   int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func NewAccount(username string) *Account {
	return &Account{
		Username: username,
	}
}
