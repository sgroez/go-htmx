package main

import (
	"fmt"
	"log"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("loading index.html")
		temp := template.Must(template.ParseFiles("index.html"))
		temp.Execute(w, nil)
	})
	fmt.Println("listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}