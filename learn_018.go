package main

import (
	"net/http"
	"log"
)

func main() {
	log.Printf("Running...")
	log.Fatal(
		http.ListenAndServe("127.0.0.1:8080",
			http.FileServer(http.Dir("/usr/share/doc"))))

}
