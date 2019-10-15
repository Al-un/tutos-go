package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/al-un/tutos-go/api_go-todos/server/router"
)

const port = 8000

func main() {
	r := router.Router()

	fmt.Printf("Starting server on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
