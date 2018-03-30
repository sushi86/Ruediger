package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print(w, "Rüdiger lebt!")
	})

	http.ListenAndServe(":80", nil)
}
