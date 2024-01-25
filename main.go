package main

import (
	"flag"
	"log"

	"github.com/Nikitalapshin95/Myproject.git/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient = telegram.New(mustToken())

}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"Токен для бота",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("Токен не указан")
	}

	return *token
}
