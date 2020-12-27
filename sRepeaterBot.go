package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const chatID = 0
const chatToken = "chatToken"
const TGI_API_TIMEOUT = 60
const IS_DEBUG_MODE = true
const SHOW_WORDS_STATISTIC_CMD = "WORD_STAT"
const SHOW_WORDS_STATISTIC_REGEXP = "[^a-zA-Z0-9]+"
const SHOW_WORDS_STATISTIC_TIMEOUT = 24

func main() {
	reg, err := regexp.Compile(SHOW_WORDS_STATISTIC_REGEXP)
	if err != nil {
		log.Fatal(err)
	}
	countWordsMap := make(map[string]int)

	var bot = initBotAPI()
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = TGI_API_TIMEOUT

	sysChannel := make(chan string)
	go func() {
		time.Sleep(SHOW_WORDS_STATISTIC_TIMEOUT * time.Hour)
		sysChannel <- SHOW_WORDS_STATISTIC_CMD
	}()

	updatesChatMsgChannel, _ := bot.GetUpdatesChan(ucfg)
	for {
		select {
		case update := <-updatesChatMsgChannel:
			activeUserName := update.Message.From.UserName
			chatID := update.Message.Chat.ID
			msgText := update.Message.Text
			log.Printf("Request content: activeUserName:%s chatID:%d msgText:%s", activeUserName, chatID, msgText)

			bot.Send(tgbotapi.NewMessage(chatID, msgText))
			for _, word := range strings.Split(reg.ReplaceAllString(msgText, ""), " ") {
				countWordsMap[word]++
			}

			if update.Message.NewChatMembers != nil && len(*update.Message.NewChatMembers) > 0 {
				newUserMessageSend(*bot, len(*update.Message.NewChatMembers), (*update.Message.NewChatMembers)[0].FirstName)
			}

		case systemUpdate := <-sysChannel:
			log.Printf("SystemUpdate was received:", systemUpdate)
			switch systemUpdate {
			case SHOW_WORDS_STATISTIC_CMD:
				showStatisticMessageSend()
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
	bot.Debug = IS_DEBUG_MODE
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return bot
}

func newUserMessageSend(bot tgbotapi.BotAPI, newChatMembersSize int, userName string) {
	log.Printf("New chat members size is:%d", newChatMembersSize)
	var msgText string = fmt.Sprintf(`Hi user: @%s! I am watching for all in here. Good luck!`, userName)
	bot.Send(tgbotapi.NewMessage(chatID, msgText))
}

func showStatisticMessageSend() {
	// TODO("not implemented")
}
