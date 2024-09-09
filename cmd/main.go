package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
