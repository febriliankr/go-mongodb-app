package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load() //load .env file
	if e != nil {
		fmt.Print(e)
	}
	fmt.Println(os.Getenv("DB_URL"))
}
