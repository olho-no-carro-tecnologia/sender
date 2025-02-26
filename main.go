package main

import (
	"log"
	"net/http"
	"os"

	"poc.sender/internal/config"
	"poc.sender/internal/dispatcher"
	"poc.sender/internal/queue"
)

// main is the entry point of the application. It performs the following tasks:
// 1. Loads environment configurations.
// 2. Connects to a message queue (RabbitMQ or SQS).
// 3. Starts a worker to process messages from the queue.
// 4. Sets up an HTTP route for a health check endpoint ("/ping").
// 5. Starts an HTTP server on the specified port (default is 8080).
func main() {
	config.LoadEnv()

	q, err := queue.Connect()
	if err != nil {
		log.Fatalf("Erro ao conectar à fila: %v", err)
	}
	defer q.Close()

	go dispatcher.StartWorker(q)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Pong! Servidor está rodando 🚀"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor rodando na porta %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
