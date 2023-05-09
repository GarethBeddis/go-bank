package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage interface {
	CreateAccount(*Account) error
	GetAccountById(int) (*Account, error)
	UpdateAccount(*Account) error
	DeleteAccount(int) error
}

type PostgresStore struct {
	db *gorm.DB
}

func (s *PostgresStore) Init() error {
	return nil
}

func (s *PostgresStore) createAccountTable() error {
	return nil
}

func (s *PostgresStore) CreateAccount(a *Account) error {
	return nil
}

func (s *PostgresStore) GetAccountById(id int) (*Account, error) {
	return nil, nil
}

func (s *PostgresStore) UpdateAccount(a *Account) error {
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	return nil
}

func NewPostgresStore() (*PostgresStore, error) {
	postgresConfig := postgres.Config{
		DSN:                  "host=localhost port=5432 user=postgres dbname=postgres password=gobank312 sslmode=disable",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}

	db, err := gorm.Open(postgres.New(postgresConfig), &gorm.Config{})

	db.AutoMigrate(&Account{})

	if err != nil {
		log.Fatalf("Cannot connect to postgres DB: %v", err)
		return nil, err
	}

	postgresDB, _ := db.DB()
	if err := postgresDB.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}
