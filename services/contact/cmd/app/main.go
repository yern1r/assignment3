package main

import (
	"fmt"
	repository "xamss/microservices/test/pkg/store/postgres"
)

func main() {
	_, err := repository.ConnPostgres("localhost", "guest", 5432, "pa55word", "mstest")
	if err != nil {
		fmt.Errorf("database connection failed: %s", err)
	}
	fmt.Println("database connection is established")
}
