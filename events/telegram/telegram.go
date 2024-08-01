package telegram

import "github.com/igortoigildin/bot-adviser/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
	// storage
}
