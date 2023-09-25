package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/jriedel-ionos/rampup-challenge-grpc/server/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PageData struct {
	Value string
}

//go:embed index.html
var templateFile embed.FS

func main() {
	if err := mainImpl(); err != nil {
		log.Fatal(err)
	}
}

func mainImpl() error {
	port := flag.Int("port", 8081, "port for the frontend")
	backendAddress := flag.String("backend", os.Getenv("TARGET"), "address of the backend server")
	flag.Parse()

	tmpl, err := template.ParseFS(templateFile, "index.html")
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	conn, err := grpc.Dial(*backendAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to backend server: %v", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("Failed to close gRPC connection: %v", err)
		}
	}()

	client := pb.NewEnvVariableClient(conn)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		varName := r.URL.Path[1:]

		resp, err := client.GetEnvironmentVariable(context.Background(), &pb.GetEnvironmentVariableRequest{
			VariableName: varName,
		})
		if err != nil {
			log.Printf("Failed to get environment variable: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		data := PageData{Value: resp.Value}

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
		return fmt.Errorf("failed to start frontend server: %v", err)
	}

	return nil
}
