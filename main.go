package main

import (
	_ "asura/commands"
	"asura/database"
	_ "asura/handler"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/servusdei2018/shards/v2"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("could not load env keys: %s", err)
	}

	conn := fmt.Sprintf("Bot %s", os.Getenv("TOKEN"))
	s, _ := shards.New(conn)

	s.AddHandler(OnReady)
	s.AddHandler(OnInteractionCreate)

	if err := database.Connect(); err != nil {
		log.Fatalf("could not connect to database: %s", err)
	}

	if err := database.Init(); err != nil {
		log.Fatalf("could not connect to redis: %s", err)
	}

	if err := s.Start(); err != nil {
		log.Fatalf("could not open session: %s", err)
	}

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sigch

	if err := s.Shutdown(); err != nil {
		log.Printf("could not close session gracefully: %s", err)
	}
}
