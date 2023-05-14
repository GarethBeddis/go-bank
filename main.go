package main

import (
	"log"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	store.db.AutoMigrate(&Account{})
	// store.db.Delete(&Account{}, "")

	// log.Printf("%+v\n", store)

	server := NewApiServer(":8080", store)
	server.Run()
}
