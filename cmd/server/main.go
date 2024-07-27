package main

import (
    "fmt"
    amqp "github.com/rabbitmq/amqp091-go"
    "os"
    "os/signal"
    "github.com/kznhq/learn-pub-sub-starter/internal/routing"
    "github.com/kznhq/learn-pub-sub-starter/internal/pubsub"
)

func main() {
	fmt.Println("Starting Peril server...")

    connection := "amqp://guest:guest@localhost:5672/"      //this is where the app will connect to the RabbitMQ server
    conn, err := amqp.Dial(connection)      //open the connection
    if err != nil {
        fmt.Errorf("Error when dialing: %v", err)
    }
    defer conn.Close()
    fmt.Println("Connection successful")

    pauseChan, err := conn.Channel()
    if err != nil {
        fmt.Errorf("Error when making channel for pause/resume messages: %v", err)
    }
    pausedState := routing.PlayingState {
        IsPaused: true,
    }
    err = pubsub.PublishJSON(pauseChan, routing.ExchangePerilDirect, routing.PauseKey, pausedState)
    if err != nil {
        fmt.Errorf("Error when pausing game: %v", err)
    }

    signalChan := make(chan os.Signal, 1)       //this Go channel will be used to wait for an interrupt signal in the terminal which will then close the server
    signal.Notify(signalChan, os.Interrupt)
    <- signalChan
    close(signalChan)
    fmt.Println("Shutting down...")
} 
