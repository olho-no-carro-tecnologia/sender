package dispatcher

import (
	"encoding/json"
	"log"

	"poc.sender/internal/email"

	"github.com/streadway/amqp"
)

// Message representa o formato das mensagens na fila
type Message struct {
	Type    string `json:"type"` // "email" ou "push"
	Payload string `json:"payload"`
}

// StartWorker consome mensagens da fila e despacha para os serviços apropriados
func StartWorker(conn *amqp.Connection) {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("[ERROR] Falha ao abrir canal: %v", err)
	}
	defer ch.Close()

	queueName := "message_queue"
	msgs, err := ch.Consume(
		queueName,
		"",
		true,  // Auto-Ack
		false, // Exclusivo
		false, // No-local
		false, // No-Wait
		nil,   // Argumentos
	)
	if err != nil {
		log.Fatalf("[ERROR] Falha ao consumir fila: %v", err)
	}

	log.Println("[INFO] Worker iniciado, aguardando mensagens...")

	for msg := range msgs {
		var message Message
		if err := json.Unmarshal(msg.Body, &message); err != nil {
			log.Printf("[WARN] Mensagem inválida: %v", err)
			continue
		}

		switch message.Type {
		case "email":
			log.Println("[INFO] Processando envio de email")
			email.SendEmail(message.Payload)
		case "push":
			log.Println("[INFO] Processando envio de push notification")
		default:
			log.Printf("[WARN] Tipo de mensagem desconhecido: %s", message.Type)
		}
	}
}
