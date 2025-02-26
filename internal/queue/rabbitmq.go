package queue

import (
	"log"
	"time"

	"poc.sender/internal/config"

	"github.com/streadway/amqp"
)

// Connect estabelece uma conexão com o RabbitMQ
func Connect() (*amqp.Connection, error) {
	url := config.GetEnv("RABBITMQ_URL", "amqp://guest:guest@rabbitmq:5672/")

	var conn *amqp.Connection
	var err error

	for i := 0; i < 5; i++ { // Tenta reconectar 5 vezes
		conn, err = amqp.Dial(url)
		if err == nil {
			log.Println("[INFO] Conectado ao RabbitMQ")
			return conn, nil
		}

		log.Printf("[WARN] Falha ao conectar ao RabbitMQ, tentativa %d: %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	return nil, err
}

// CreateChannel cria um canal e declara a fila dentro da conexão RabbitMQ
func CreateChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	// Declara a fila (cria se não existir)
	_, err = ch.QueueDeclare(
		"message_queue", // Nome da fila
		true,            // Persistente
		false,           // Não deletar automaticamente
		false,           // Exclusiva
		false,           // Sem espera
		nil,             // Argumentos extras
	)
	if err != nil {
		return nil, err
	}

	log.Println("[INFO] Fila 'message_queue' garantida no RabbitMQ")
	return ch, nil
}
