package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.RedirectHandler(os.Getenv("REDIRECT_URL"), 301))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
