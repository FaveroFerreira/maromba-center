package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/faveroferreira/maromba-center/internal"
	"github.com/faveroferreira/maromba-center/internal/configs"
)

func main() {
	profile := os.Getenv("PROFILE")
	if profile == "" {
		profile = "local"
	}

	err := configs.LoadConfig(profile)
	if err != nil {
		panic(err)
	}

	ctrlC := make(chan os.Signal, 1)
	signal.Notify(ctrlC, os.Interrupt)

	go internal.StartGin(profile)

	signal := <-ctrlC

	fmt.Println("Application exited: ", signal)
}
