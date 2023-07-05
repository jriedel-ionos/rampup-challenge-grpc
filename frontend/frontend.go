package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
)

type PageData struct {
	Value string
}

//go:embed index.html
var templateFile embed.FS

func main() {
	const Port = "8081"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		varName := r.URL.Path[1:]
		value := os.Getenv(varName)

		data := PageData{Value: value}

		tmpl, err := template.ParseFS(templateFile, "index.html")
		if err != nil {
			log.Fatalf("Failed to parse template: %v", err)
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println("Failed to render template:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	log.Println("Frontend server started on port " + Port)
	log.Fatal(http.ListenAndServe(":"+Port, nil))
}
