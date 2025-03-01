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
	validPaths := map[string]bool{
		"": true,
		"editor": true,
		"images": true,
	}
	//check if path is valid
	if _, contains := validPaths[path]; !contains {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
	files := []string{"index.html"}
	//if path not empty add subpage template to files
	if len(path) > 0 {
		files = append(files, path + ".html")
	}
	//load template files
	tmpl := template.Must(template.ParseFiles(files...))
	//calculate data for the templates and execute the templates with the dat
	switch path {
	case "images":
		fileNames, err := extractFileNames("./static/images")
		if err != nil {
			http.Error(w, "error while loading images", http.StatusInternalServerError)
		}
		tmpl.Execute(w, fileNames)
	default:
		tmpl.Execute(w, nil)
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