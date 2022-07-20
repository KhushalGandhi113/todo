package main

import (
	"fmt"
	"todo/routers"
)

func main() {
	router := routers.RegisterRoutes()
	fmt.Printf("\nSuccessfully connected to database!\n")

	router.Run("localhost:8080")
}
