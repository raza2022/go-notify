package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe(":3001", nil)
}

//reference https://www.alexedwards.net/blog/serving-static-sites-with-go
