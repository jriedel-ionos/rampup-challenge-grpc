package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type PageData struct {
	Value string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		varName := r.URL.Path[1:]
		value := os.Getenv(varName)

		data := PageData{Value: value}

		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			log.Println("Failed to parse template:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println("Failed to render template:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	log.Println("Frontend server started on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
