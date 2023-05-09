package main

import (
	"math/rand"

	_ "github.com/lib/pq"
)

type Account struct {
	ID       int    `json:"id"` // todo: change to uuid
	Username string `json:"username"`
	Number   int64  `json:"number"`
	Balance  int64  `json:"balance"`
}

func NewAccount(username string) *Account {
	return &Account{
		ID:       rand.Intn(100000),
		Username: username,
		Number:   int64(rand.Intn(10000000)),
	}
}
