package handler

import (
	"github.com/hudson6666/nanami/haruka"
	"github.com/hudson6666/nanami/hoshino"
	"time"
	"strings"
	log "github.com/Sirupsen/logrus"
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
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
			if m.IsCommand() && (m.Chat.IsPrivate() || strings.Contains(m.Text, "@nanami_nanabot")) {
				log.Infof("Chat ID: %d", m.Chat.ID)
				t := haruka.HandleCommand(m.Command(), m.CommandArguments(), m.From.ID, m.Chat.ID)
				replyMessage(t, bot, m)
			} else {
				log.Infof("Chat ID: %d", m.Chat.ID)
				if r, t := haruka.HandleText(m.Text, m.From.ID); r {
					replyMessage(t, bot, m)
				}
			}
		case "Hoshino":
			if m.IsCommand() && (m.Chat.IsPrivate() || strings.Contains(m.Text, "@nanami_alphabot")) {
				log.Infof("Chat ID: %d", m.Chat.ID)
				t := hoshino.HandleCommand(m.Command(), m.CommandArguments())
				replyMessage(t, bot, m)
			} else {
				log.Infof("Chat ID: %d", m.Chat.ID)
			}
		}
	}

}

func replyMessage(text string, bot *tg.BotAPI, req *tg.Message) {
	msg := tg.NewMessage(req.Chat.ID, text)
	msg.ReplyToMessageID = req.MessageID
	bot.Send(msg)
}