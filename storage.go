package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage interface {
	CreateAccount(*Account) error
	GetAccountById(int) (*Account, error)
	GetAccounts() ([]*Account, error)
	UpdateAccount(*Account) error
	DeleteAccount(int) error
}

type PostgresStore struct {
	db *gorm.DB
}

func (s *PostgresStore) Init() error {
	return s.db.AutoMigrate(&Account{})
}

func (s *PostgresStore) CreateAccount(a *Account) error {
	result := s.db.Create(&a)

	if err := result.Error; err != nil {
		return err
	}

	fmt.Printf("%+v\n", result)

	return nil
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	var accounts []*Account

	res := s.db.Find(&accounts)

	if res.Error != nil {
		fmt.Println(res.Error)
	}

	return accounts, nil
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
