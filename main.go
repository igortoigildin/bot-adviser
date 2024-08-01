package main

import (
	"flag"
	"log"

	"github.com/igortoigildin/bot-adviser/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {

	_ = telegram.New(tgBotHost, mustToken())
}

func mustToken() string {
	token := flag.String("token-bot-token", "", "token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
