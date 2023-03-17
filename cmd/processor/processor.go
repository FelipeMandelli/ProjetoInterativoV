package main

import (
	"log"

	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	zapConfig := zap.NewProductionConfig()

	zapConfig.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	logger, err := zapConfig.Build()
	if err != nil {
		log.Fatal("error creating logger")
	}

	logger.Info("This is the processor application!")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Não foi possível estabelecer a conexão com o RabbitMQ: %s", err)
	}
	defer conn.Close()

	// Cria um canal de comunicação
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Não foi possível criar o canal: %s", err)
	}
	defer ch.Close()

	// Declara a fila a ser consumida
	q, err := ch.QueueDeclare(
		"minha-fila", // Nome da fila
		false,        // Durable
		false,        // Delete when unused
		false,        // Exclusive
		false,        // No-wait
		nil,          // Arguments
	)
	if err != nil {
		log.Fatalf("Não foi possível declarar a fila: %s", err)
	}

	// Registra um consumidor
	msgs, err := ch.Consume(
		q.Name, // Nome da fila
		"",     // Consumer
		true,   // Auto-ack
		false,  // Exclusive
		false,  // No-local
		false,  // No-wait
		nil,    // Arguments
	)
	if err != nil {
		log.Fatalf("Não foi possível registrar o consumidor: %s", err)
	}

	// Escuta por mensagens
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Recebida mensagem: %s", d.Body)
		}
	}()
	logger.Info("Escutando por mensagens. Para sair pressione CTRL+C")
	<-forever
}
