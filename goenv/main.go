package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = godotenv.Load(".env.dev")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println(os.Getenv("APP_ENV"))
	fmt.Println(os.Getenv("MYSQL_PORT"))
}
