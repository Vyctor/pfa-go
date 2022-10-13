package rabbitmq

import "github.com/streadway/amqp"

func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	err = ch.Qos(100, 0, false)

	if err != nil {
		panic(err)
	}

	return ch, nil
}

func Consume(channel *amqp.Channel, output chan amqp.Delivery) error {
	messages, err := channel.Consume("orders", "go-consumer", false, false, false, false, nil)

	if err != nil {
		return err
	}

	for message := range messages {
		output <- message
	}

	return nil
}
