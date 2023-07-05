package main

import (
	"embed"
	"flag"
	"fmt"
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
	port := flag.Int("port", 8081, "port for the frontend")
	flag.Parse()

	tmpl, err := template.ParseFS(templateFile, "index.html")
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		varName := r.URL.Path[1:]
		value := os.Getenv(varName)

		data := PageData{Value: value}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Printf("Failed to render template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	log.Printf("Frontend server started on port %v", *port)
	err = http.ListenAndServe(fmt.Sprintf(":%v", *port), nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
