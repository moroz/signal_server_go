package main

import (
	"fmt"
	"os"
)

func main() {
	database_url := os.Getenv("DATABASE_URL")
	fmt.Println(database_url)
}
