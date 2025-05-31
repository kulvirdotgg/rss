package main

import (
	"fmt"
	"os"

	env "github.com/joho/godotenv"
)

func main() {
	env.Load()

	port := os.Getenv("PORT")
	fmt.Println(port)
}
