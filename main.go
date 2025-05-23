package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file") 
	} 

	// apiKey := os.Getenv("CMC_KEY")


	fmt.Println("Hello World")
}

