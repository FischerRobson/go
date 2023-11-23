package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"imersao-fullcycle/internal/infra/database"
	"imersao-fullcycle/internal/usecase"
	"imersao-fullcycle/pkg/rabbitmq"

	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}

	uc := &usecase.CalculateFinalPrice{
		OrderRepository: &database.OrderRepository{
			Db: db,
		},
	}

	input := usecase.OrderInput{
		ID:    "7",
		Price: 10,
		Tax:   1,
	}

	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", *output)

	defer db.Close()

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()
	msgRabbitmqChannel := make(chan amqp.Delivery)
	go rabbitmq.Consumer(ch, msgRabbitmqChannel)
	rabbitmqWorker(msgRabbitmqChannel, uc)
}

func rabbitmqWorker(msgChan chan amqp.Delivery, uc *usecase.CalculateFinalPrice) {
	fmt.Println("Starting rabbitmq worker")
	for msg := range msgChan {
		var input usecase.OrderInput
		err := json.Unmarshal(msg.Body, &input)
		if err != nil {
			panic(err)
		}

		output, err := uc.Execute(input)
		if err != nil {
			panic(err)
		}
		msg.Ack(false)
		fmt.Println(output)
	}
}
