package main 

import (
	"net/http"
	"fmt"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Backend hello world!")
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}