package email

import (
	"log"

	"poc.sender/internal/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

// SendEmail envia um e-mail via Amazon SES
func SendEmail(payload string) {
	awsRegion := config.GetEnv("AWS_REGION", "us-east-1")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
	})
	if err != nil {
		log.Printf("[ERROR] Falha ao iniciar sessão AWS: %v", err)
		return
	}

	svc := ses.New(sess)

	input := &ses.SendEmailInput{
		Source: aws.String(config.GetEnv("SES_EMAIL_FROM", "no-reply@meu-projeto.com")),
		Destination: &ses.Destination{
			ToAddresses: []*string{aws.String("destinatario@example.com")},
		},
		Message: &ses.Message{
			Subject: &ses.Content{
				Data: aws.String("Assunto do Email"),
			},
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String(payload),
				},
			},
		},
	}

	_, err = svc.SendEmail(input)
	if err != nil {
		log.Printf("[ERROR] Falha ao enviar email: %v", err)
		return
	}

	log.Println("[INFO] Email enviado com sucesso!")
}
