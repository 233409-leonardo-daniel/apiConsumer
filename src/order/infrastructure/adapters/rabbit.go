// src/order/infrastructure/adapters/rabbit_repository.go
package adapters

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQRepository struct{}

func NewRabbitMQRepository() *RabbitMQRepository {
	return &RabbitMQRepository{}
}

func (r *RabbitMQRepository) Publish(idProduct int32, quantity int32, totalPrice float64, status string) error {
	conn, err := amqp.Dial("amqp://leo:1234@34.235.202.211:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"order",  // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := bodyFrom(os.Args)
	err = ch.PublishWithContext(ctx,
		"order", // exchange
		"1234",  // routing key
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		return err
	}

	log.Printf(" [x] Sent %s", body)
	return nil
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "pedido recibido"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
