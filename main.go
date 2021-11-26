package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/titolins/ecr_cleaner/cmd"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cmd.Execute()
}
