package main

import (
	"fmt"
	"log"

	"github.com/dzacca/go/http_server_golang/internal/database"
)

func main() {
	c := database.NewClient("db.json")
	err := c.EnsureDB()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("database ensured!")
}