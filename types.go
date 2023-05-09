package main

import (
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Account struct {
	ID        uint32    `json:"id"`
	Number    uint32    `json:"number"`
	Username  string    `json:"username"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateAccount(username string) *Account {
	return &Account{
		ID:        uuid.New().ID(),
		Number:    uuid.New().ID(),
		Username:  username,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}

type CreateAccountRequest struct {
	Username string `json:"username"`
}
