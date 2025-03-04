package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleSiteLoad)
	http.HandleFunc("/navigate/", handleRouteUpdate)
	http.HandleFunc("/api/", handleApiCall)

	//serve static files
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/static/favicon.ico", http.StatusMovedPermanently)
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}