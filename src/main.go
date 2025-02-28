package main

import (
	"fmt"
	"log"
	"io"
	"html/template"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, nil)
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request){
		bodyByte, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "an error occured while processing the request", http.StatusInternalServerError)
			return
		}
		bodyString := string(bodyByte[:])
		bodyString, found := strings.CutPrefix(bodyString, "editor=")
		if !found {
			fmt.Println("did not find expected body string prefix")
			http.Error(w, "an error occured while processing the request", http.StatusInternalServerError)
		}
		fmt.Println(bodyString)
		tmpl := template.Must(template.ParseFiles("output.html"))
		tmpl.Execute(w, bodyString)
	})
	fmt.Println("listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}