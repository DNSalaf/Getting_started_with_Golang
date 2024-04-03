package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	path_len := len(r.URL.Path)

	println("Request received")

	if path_len == 1 {
		fmt.Fprintf(w, "Why don't you try adding a path to my url? 🤔")
	} else {
		fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	}

}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
