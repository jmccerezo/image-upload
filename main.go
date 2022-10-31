package main

import (
	"fmt"
	"net/http"

	"github.com/jmccerezo/image-upload/views"
)

func main() {
	fmt.Println("HELLO WORLD")
	fmt.Println("Go to http://localhost:8080/")
	http.HandleFunc("/", views.IndexHandler)
	http.ListenAndServe(":8080", nil)
}
