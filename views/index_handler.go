package views

import (
	"fmt"
	"net/http"
	"path"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	filepath, err := ImageUpload(r, "file")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	if filepath != "" {
		fmt.Println("FILEPATH: ", filepath)
	}
}
