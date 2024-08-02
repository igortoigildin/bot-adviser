package main

import (
	"flag"
	"log"

	tgClient "github.com/igortoigildin/bot-adviser/clients/telegram"
	event_consumer "github.com/igortoigildin/bot-adviser/consumer/event-consumer"
	"github.com/igortoigildin/bot-adviser/events/telegram"
	"github.com/igortoigildin/bot-adviser/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

func main() {
	eventsProcessor := telegram.New(tgClient.New(tgBotHost, mustToken()), files.New(storagePath))

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String("tg-bot-token", "", "token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
