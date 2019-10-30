package main

import (
	"donosbot/donos"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strings"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	bot, err := tgbotapi.NewBotAPI(viper.GetString("bot.telegramAPIKey"))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(viper.GetString("bot.botListenDomain") + bot.Token))
	if err != nil {
		log.Fatal(err)
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}
	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServe(viper.GetString("bot.botListenAndServe"), nil)

	for update := range updates {
		if strings.Contains(update.Message.Text, "/donos") {
			go donos.SendDonosMessage(bot, update)
		} else if strings.Contains(update.Message.Text, "/help") {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, `Напишите <code>/donos Текст доноса</code>
Дежурный сотрудник рассмотрит ваше обращение в ближайшее время. Спасибо.`)
			msg.ParseMode = "HTML"
			bot.Send(msg)
		}
	}
}
