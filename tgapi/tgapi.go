package tgapi

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TgApi interface {
	SendMessage(msg string)
}

type tgApi struct {
	bot       *tgbotapi.BotAPI
	targetUid string
	chatId    int64
}

func (t *tgApi) SendMessage(msg string) {
	m := tgbotapi.NewMessage(t.chatId, msg)
	t.bot.Send(m)
}

func New(Token, TargetUid string) TgApi {
	if Token == "" || TargetUid == "" {
		panic("Set target username and bot token in /etc/pgn/pgn.conf")
	}
	t := &tgApi{
		bot:       nil,
		targetUid: TargetUid,
		chatId:    -1,
	}
	bot, err := tgbotapi.NewBotAPI(Token)
	if err != nil {
		panic(err)
	}
	t.bot = bot
	c := tgbotapi.UpdateConfig{}
	u, err := t.bot.GetUpdates(c)
	if err != nil {
		panic(err)
	}
	for _, up := range u {
		if up.Message.Chat.UserName == TargetUid {
			t.chatId = up.Message.Chat.ID
			break
		}
	}

	if t.chatId == -1 {
		panic("Target UID not found! Make sure you started chat with this bot!")
	}

	return t
}
