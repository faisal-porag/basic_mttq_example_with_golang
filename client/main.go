package main

import (
	"basic_mqtt_example_with_golang/mqtt_client"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	client := mqtt_client.NewClient(mqtt_client.BrokerAddress, mqtt_client.ClientID)
	if err := client.ConnectClient(); err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		for {
			latitude := generateRandomLocation()
			longitude := generateRandomLocation()

			payload := fmt.Sprintf("{\"latitude\": %f, \"longitude\": %f}", latitude, longitude)
			if err := client.Publish(mqtt_client.Topic, payload); err != nil {
				fmt.Println("Error publishing:", err)
			}

			fmt.Printf("payload: %v", payload)
			fmt.Println("")

			time.Sleep(10 * time.Second)
		}
	}()

	<-c
}

func generateRandomLocation() float64 {
	return 40 + (rand.Float64() * 0.1) - 0.05
}
