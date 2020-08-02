package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// handle '/' route
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello HTTPS word ! here I GO!")
	})

	//run server on port "8710"
	log.Fatal(http.ListenAndServe(":8710", nil))
}
