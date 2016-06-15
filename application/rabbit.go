package application

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
)

type Rabbit struct {
	Conn       *amqp.Connection
	ActionChan *amqp.Channel
	Queue      *amqp.Queue
}

func NewRabbit(c *Config) *Rabbit {

	return &Rabbit{}
}

func (r *Rabbit) SendAction(msg string) error {
	fmt.Println("send message: [%s]", msg)
	conn, err := amqp.Dial(os.Getenv("CLOUDAMQP_URL"))
	if err != nil {
		log.Println("cannot connect to rabbitmq")
		return err
	}
	ch, err := conn.Channel()
	defer ch.Close()
	if err != nil {
		log.Println("cannot open channel: %s", err)
		return err
	}
	err = ch.ExchangeDeclare(
		"action2", // name
		"fanout",  // type
		false,     // durable
		false,     // auto-deleted
		false,     // internal
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Println("cannot exchange declare")
		log.Println(err)
		return err
	}
	err = ch.Publish(
		"action2", // exchange
		"",        // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *Rabbit) SendProgram(msg []byte) error {
	fmt.Println("send program: [%s]", msg)
	conn, err := amqp.Dial(os.Getenv("CLOUDAMQP_URL"))
	if err != nil {
		log.Println("cannot connect to rabbitmq")
		return err
	}
	ch, err := conn.Channel()
	defer ch.Close()
	if err != nil {
		log.Println("cannot open channel: %s", err)
		return err
	}
	err = ch.ExchangeDeclare(
		"program", // name
		"fanout",  // type
		false,      // durable
		false,     // auto-deleted
		false,     // internal
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Println("cannot exchange declare")
		log.Println(err)
		return err
	}
	err = ch.Publish(
		"program", // exchange
		"",       // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(msg),
		})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
