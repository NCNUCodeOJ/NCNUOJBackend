package submission

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// 傳送的資料結構
type Answer = struct {
	ProblemID    string `json:"problemid"`
	Code         string `json:"code"`
	CodeLanguage string `json:"codelanguage"`
}

func Send_failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Send_main(data []byte) {
	conn, err := amqp.Dial("amqp://root:admin1234@localhost:5672/")
	Send_failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	Send_failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	Send_failOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         data,
		})
	Send_failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", string(data))
}

func BodyFrom(answer Answer) []byte {
	if s, err := json.Marshal(answer); err == nil {
		return s
	} else {
		return []byte("")
	}

}
