package main

import (
	"bot/models"
	"bot/pkg"
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	//pkg.Analogi("Toyota", "90915yzze2")
	bot, err := tgbotapi.NewBotAPI("2018104273:AAEvHzqS3MX9-qei0lnhaXiG5iqS-d6ZmKg")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		fmt.Printf("update_message = %v\n", update.Message.Text)
		a := pkg.SqlReq(update.Message.Text)
		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		message := "Мытищи: " + strconv.Itoa(a[0].Qtym) + " штук " + a[0].Cellm + "\n Титан: " + strconv.Itoa(a[0].Qtyt) + " штук" + a[0].Cellt + "\n Цена " + a[0].Price + " руб\n"
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, message) // update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID
		log.Printf("msg=%v\n", msg)
		bot.Send(msg)
		pkg.Analogi(a[0].Firm, a[0].Partnum)
		for _, analog := range models.Analogs {
			a = pkg.SqlReq(analog.Number)
			fmt.Printf("очередное =%v a=%v\n", analog, a)
			if a[0].Qtyt != 0 || a[0].Qtym != 0 {
				fmt.Printf("Такой аналог есть\n")
				message := analog.Firm + "  " + analog.Number + "\nМытищи: " + strconv.Itoa(a[0].Qtym) + " штук " + a[0].Cellm + "  Титан: " + strconv.Itoa(a[0].Qtyt) + " штук" + a[0].Cellt + "\n Цена " + a[0].Price + " руб\n"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, message) // update.Message.Text)
				msg.ReplyToMessageID = update.Message.MessageID
				log.Printf("msg=%v\n", msg)
				bot.Send(msg)
			}

		}
	}
}

