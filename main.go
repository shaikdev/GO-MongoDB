package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shaikdev/GO-MongoDB/routes"
)

func main() {
	r := routes.Router()
	fmt.Println("Server running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", r))

}
