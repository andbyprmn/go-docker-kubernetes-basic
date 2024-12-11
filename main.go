package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT env is required")
	}

	instanceID := os.Getenv("INSTANCE_ID")

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "The requested HTTP Method is not allowed", http.StatusMethodNotAllowed)
			return
		}

		text := "Hello world"
		if instanceID != "" {
			text = text + ". from " + instanceID
		}

		w.Write([]byte(text))
	})

	server := new(http.Server)
	server.Handler = mux
	server.Addr = ":" + port

	log.Println("Web server is starting at", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
