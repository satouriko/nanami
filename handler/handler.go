package handler

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/hudson6666/nanami/haruka"
	"time"
	log "github.com/Sirupsen/logrus"
	"github.com/hudson6666/nanami/hoshino"
)

func Init()  {
	//只有一个 Nanami 不要冒充哦
	go serveTelegram("Haruka", "419149697:AAG_x2ITdgzk4NWr6vAQ1YpAdkeMJp5SvLw")
	go serveTelegram("Hoshino", "259874633:AAEz2q9K9rDnR4xDEXHrzTa4KFGl9nt92og")
}

func serveTelegram(botVersion string, apiKey string)  {
	log.Infof("Nanami %s Started at %s", botVersion, time.Now())
	bot, err := tg.NewBotAPI(apiKey)
	if err != nil {
		log.Fatal(err)
	}
	u := tg.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		m := update.Message

		switch botVersion {
		case "Haruka":
			if m.IsCommand() {
				log.Infof("Chat ID: %d", m.Chat.ID)
				t := haruka.HandleCommand(m.Command(), m.CommandArguments())
				replyMessage(t, bot, m)
			} else {
				continue
			}
		case "Hoshino":
			if m.IsCommand() {
				log.Infof("Chat ID: %d", m.Chat.ID)
				t := hoshino.HandleCommand(m.Command(), m.CommandArguments())
				replyMessage(t, bot, m)
			} else {
				log.Infof("Chat ID: %d", m.Chat.ID)
				t := hoshino.HandleText(m.Text)
				replyMessage(t, bot, m)
			}
		}
	}

}

func replyMessage(text string, bot *tg.BotAPI, req *tg.Message) {
	msg := tg.NewMessage(req.Chat.ID, text)
	msg.ReplyToMessageID = req.MessageID
	bot.Send(msg)
}