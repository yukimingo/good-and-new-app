package main

import (
	"good-and-new/internal/db"
	"good-and-new/internal/domain"
	"log"
)

func main() {
	db := db.NewDB()

	if err := db.AutoMigrate(domain.User{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}
}
