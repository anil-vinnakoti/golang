package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anil-vinnakoti/mongoAPI/router"
)

func main() {
	r := router.Router()
	fmt.Println("Server is getting started...")

	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("listening at port 4000")
}
