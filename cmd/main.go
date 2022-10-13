package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/vyctor/pfa-go/internal/order/infra/database"
	"github.com/vyctor/pfa-go/internal/order/usecase"
	"github.com/vyctor/pfa-go/pkg/rabbitmq"
	"net/http"
	"sync"
)

func main() {
	maxWorkers := 10

	waitGroup := sync.WaitGroup{}

	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	repository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPriceUseCase(repository)

	http.HandleFunc("/total", func(w http.ResponseWriter, r *http.Request) {
		uc := usecase.NewGetTotalUseCase(repository)
		output, err := uc.Execute()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(output)
	})

	go http.ListenAndServe(":8181", nil)

	channel, err := rabbitmq.OpenChannel()

	if err != nil {
		panic(err)
	}

	defer channel.Close()

	out := make(chan amqp.Delivery)
	go rabbitmq.Consume(channel, out)

	waitGroup.Add(maxWorkers)

	for i := 0; i < maxWorkers; i++ {
		fmt.Println("Starting worker: ", i)
		defer waitGroup.Done()
		go worker(out, uc, i)
	}
	waitGroup.Wait()
}

func worker(deliveryMessage <-chan amqp.Delivery, uc *usecase.CalculateFinalPriceUseCase, workerId int) {
	for message := range deliveryMessage {
		var input usecase.OrderInputDTO
		err := json.Unmarshal(message.Body, &input)
		if err != nil {
			fmt.Println("Error unmarshalling message:", err)
		}
		input.Tax = 10
		_, err = uc.Execute(input)
		if err != nil {
			fmt.Println("Error ack message:", err)
		}
		err = message.Ack(false)
		fmt.Println("Worker: ", workerId, " - Message processed", input.ID)
	}
}
