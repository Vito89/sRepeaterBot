package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const isDebugMode = true
const chatID = 0
const chatToken = "chatToken"
const tgiAPITimeout = 60

func main() {
	var bot = initBotAPI()
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = tgiAPITimeout

	updatesChannel, _ := bot.GetUpdatesChan(ucfg)
	for {
		select {
		case update := <-updatesChannel:
			activeUserName := update.Message.From.UserName
			chatID := update.Message.Chat.ID
			msgText := update.Message.Text
			log.Printf("Request content: activeUserName:%s chatID:%d msgText:%s", activeUserName, chatID, msgText)

			bot.Send(tgbotapi.NewMessage(chatID, msgText))

			if update.Message.NewChatMembers != nil && len(*update.Message.NewChatMembers) > 0 {
				newUserMessageSend(*bot, len(*update.Message.NewChatMembers), (*update.Message.NewChatMembers)[0].FirstName)
			}
		}
	}
}

func initBotAPI() *tgbotapi.BotAPI {
	log.Printf("Sending tgBotAPI.NewBotAPI POST")

	bot, err := tgbotapi.NewBotAPI(fmt.Sprintf("%s:%s", chatID, chatToken))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = isDebugMode
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return bot
}

func newUserMessageSend(bot tgbotapi.BotAPI, newChatMembersSize int, userName string) {
	log.Printf("New chat members size is:%d", newChatMembersSize)
	var msgText string = fmt.Sprintf(`Hi user: @%s! I am watching for all in here. Good luck!`, userName)
	bot.Send(tgbotapi.NewMessage(chatID, msgText))
}
