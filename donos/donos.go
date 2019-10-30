package donos

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"math/rand"
	"time"
)

// SendTyping - Send status to chat for humanity
func SendTyping(b *tgbotapi.BotAPI, m tgbotapi.Update, t <-chan bool) bool {
	for {
		select {
		case a := <-t:
			return a
		default:
			_, err := b.Send(tgbotapi.NewChatAction(m.Message.Chat.ID, tgbotapi.ChatTyping))
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func RandomTimeWait(t chan<- bool) {
	rand.Seed(time.Now().UnixNano())
	min := 4
	max := 10
	wt := rand.Intn(max-min+1) + min
	time.Sleep(time.Second * time.Duration(wt))
	t <- true
}

func DonosReceived(b *tgbotapi.BotAPI, m tgbotapi.Update) {
	msg := tgbotapi.NewMessage(m.Message.Chat.ID, "Ваш донос принят")
	msg.ReplyToMessageID = m.Message.MessageID
	b.Send(msg)
}

func SendSticker(b *tgbotapi.BotAPI, m tgbotapi.Update) {
	msg := tgbotapi.NewStickerShare(m.Message.Chat.ID, "CAADBAADCgMAAlGMzwFUU5UxUx_P9xYE")
	msg.ReplyToMessageID = m.Message.MessageID
	b.Send(msg)
}

func SendDonosMessage(b *tgbotapi.BotAPI, m tgbotapi.Update) {
	t := make(chan bool)
	s := make(chan bool)
	go RandomTimeWait(t)
	SendTyping(b, m, t)
	DonosReceived(b, m)

	go RandomTimeWait(s)
	SendTyping(b, m, s)
	SendSticker(b, m)
}
