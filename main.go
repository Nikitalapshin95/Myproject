package main

import (
	"flag"
	"log"

	tgClient "github.com/Nikitalapshin95/Myproject.git/clients/telegram"
	event_consumer "github.com/Nikitalapshin95/Myproject.git/consumer/event-consumer"
	"github.com/Nikitalapshin95/Myproject.git/events/telegram"
	"github.com/Nikitalapshin95/Myproject.git/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"token",
		"",
		"Токен для бота",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("Токен не указан")
	}

	return *token
}
