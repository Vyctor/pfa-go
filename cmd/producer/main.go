package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"math/rand"
)

type Order struct {
	ID    string
	Price float64
}

func GenerateOrders() Order {
	return Order{
		ID:    uuid.New().String(),
		Price: rand.Float64()*1000 + 7,
	}
}

func Notify(channel *amqp.Channel, order Order) error {
	body, err := json.Marshal(order)
	if err != nil {
		return err
	}
	err = channel.Publish(
		"amq.direct",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	return err
}

func main() {
	connection, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		panic(err)
	}

	defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	for i := 0; i < 2000000; i++ {
		order := GenerateOrders()
		err := Notify(channel, order)
		if err != nil {
			panic(err)
		}
		fmt.Println("Order created: ", order)
	}
}
