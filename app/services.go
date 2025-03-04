package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"net/http"
	"html/template"
)

func handleSiteLoad(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if isValid := validatePath(path); !isValid {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
	files := []string{"index.html"}
	files = append(files, path + ".html")
	tmpl := template.Must(template.ParseFiles(files...))
	data, err := loadRouteData(path)
	if err != nil {
		http.Error(w, "error while loading data", http.StatusInternalServerError)
	}
	tmpl.Execute(w, data)
}

func handleRouteUpdate(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[10:]
	if isValid := validatePath(path); !isValid {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
	tmpl := template.Must(template.ParseFiles(path + ".html"))
	data, err := loadRouteData(path)
	if err != nil {
		http.Error(w, "error while loading data", http.StatusInternalServerError)
	}
	tmpl.Execute(w, data)
}

func validatePath(path string) bool {
	validPaths := map[string]bool{
		"home": true,
		"editor": true,
		"images": true,
	}
	if _, contains := validPaths[path]; contains {
		return true
	}
	return false
}

func loadRouteData(path string) (any, error) {
	switch path {
	case "images":
		fileNames, err := extractFileNames("./static/images")
		if err != nil {
			return nil, err
		}
		return fileNames, nil
	default:
		return nil, nil
	}
}

func handleApiCall(w http.ResponseWriter, r *http.Request) {
	body := readRequestBodyString(w, r)
	tmpl := template.Must(template.ParseFiles("output.html"))
	tmpl.Execute(w, body)
}

func extractFileNames(folderPath string) ([]string, error) {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}
	fileNames := []string{}
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	return fileNames, nil
}

func readRequestBodyString(w http.ResponseWriter, r *http.Request) string {
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "an error occured while processing the request", http.StatusInternalServerError)
	}
	bodyString := string(bodyByte[:])
	bodyString, found := strings.CutPrefix(bodyString, "editor=")
	if !found {
		fmt.Println("did not find expected body string prefix")
		http.Error(w, "an error occured while processing the request", http.StatusInternalServerError)
	}
	return bodyString
}