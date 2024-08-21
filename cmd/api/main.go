package main

import (
	"fmt"
	"go-gin-api/internal/database"
	"go-gin-api/internal/server"
)

func main() {
	fmt.Println("Hello go gin api")

	database.Connect();

	ser := server.NewServer()

	// Print the server address before starting
	fmt.Println("Server address:", ser.Addr)

	err := ser.ListenAndServe()

	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}