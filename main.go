package main

import (
	"log"
	"net/http"
	"os"

	"easy-ai.us/openai/v1/proxy"
)

func main() {

	http.HandleFunc("/v1/health", func(w http.ResponseWriter, r *http.Request) {
		log.Print("health check...")
		w.Write([]byte("ok"))
	})

	proxy := proxy.NewOpenAIProxy()

	http.HandleFunc("/v1/chat/completions", func(rw http.ResponseWriter, req *http.Request) {
		proxy.ServeHTTP(rw, req)
	})

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "80"
	}

	log.Print("Server started")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
