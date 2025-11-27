package main

import (
	"context"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Artifact struct {
	gorm.Model
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {
	DATABASE_URL := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	ctx := context.Background()

	// Migrate the schema
	db.AutoMigrate(&Artifact{})

	// // Create
	// error := gorm.G[Artifact](db).Create(ctx, &Artifact{ID: 2, Title: "Blue Lagoon", Artist: "Joe Biden", Price: 56.99})

	// if error != nil {
	// 	fmt.Println("Error creating album:", error)
	// }

	// Read
	artifact, err := gorm.G[Artifact](db).Where("id = ?", 2).First(ctx)
	if err != nil {
		fmt.Println("Error finding artifact:", err)
	}
	artifacts, error := gorm.G[Artifact](db).Find(ctx)
	if error != nil {
		fmt.Println("Error finding artifacts:", error)
	}
	fmt.Println(artifact)
	fmt.Println(artifacts)
}
