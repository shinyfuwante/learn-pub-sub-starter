package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril server...")
	conn_string := "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(conn_string)
	if err == nil {
		fmt.Println(("Connected to RabbitMQ server"))
	} else {
		fmt.Println("Failed to connect to RabbitMQ server")
		return
	}
	defer conn.Close()
	channel, err := conn.Channel()
	if err != nil {
		fmt.Println("Failed to create pause/resume channel")
		return
	}
	pubsub.PublishJSON(channel, routing.ExchangePerilDirect, routing.PauseKey, json.Marshal(routing.PlayingState{IsPaused: true}))

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	fmt.Println("Stopping Peril server...")

}
