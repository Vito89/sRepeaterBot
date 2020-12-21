package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func main() {
	log.Printf("Sending tgBotAPI.NewBotAPI POST")

	bot, err := tgbotapi.NewBotAPI("id:key")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	upd, err := bot.GetUpdatesChan(ucfg)
	for {
		select {
		case update := <-upd:
			activeUserName := update.Message.From.UserName
			chatID := update.Message.Chat.ID
			msgText := update.Message.Text
			log.Printf("Request content: activeUserName:%s chatID:%d msgText:%s", activeUserName, chatID, msgText)

			bot.Send(tgbotapi.NewMessage(chatID, msgText))

			if update.Message.NewChatMembers != nil && len(*update.Message.NewChatMembers) > 0 {
				log.Printf("New chat members size is:%d", len(*update.Message.NewChatMembers))
				var reply string = fmt.Sprintf(`Hi user: @%s! I am watching for all in here. Good luck!`,
					(*update.Message.NewChatMembers)[0].FirstName)
				bot.Send(tgbotapi.NewMessage(chatID, reply))
			}
		}
	}
}
