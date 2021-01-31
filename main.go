package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	users    = make(map[string]int64) // [username]chatId
	messages = make(map[int64]int)    // [chatId]messageId
	sh_list  []string
	bot      *tgbotapi.BotAPI
)

func main() {
	var err error
	bot, err = tgbotapi.NewBotAPI("1543629416:AAH3atg50vYtHA5VhgEXdIF3uFXGCsqlGSU")
	if err != nil {
		log.Panic(err)
	}

	// if err := database.Connect("127.0.0.1:27017", "shlist"); err != nil {
	// 	return
	// }
	// defer database.Disconnect()

	bot.Debug = false

	// log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			if update.Message.From.UserName != "vitfil" && update.Message.From.UserName != "Julie1908" {
				bot.DeleteMessage(tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID))
				continue
			}
			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case "start":
					users[update.Message.From.UserName] = update.Message.Chat.ID
					sendList(update.Message.Chat.ID)
					bot.DeleteMessage(tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID))
				}
			} else {
				bot.DeleteMessage(tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID))
				sh_list = append(sh_list, update.Message.Text)
				updateList()
			}
		} else if update.CallbackQuery != nil {
			cmd := strings.Split(update.CallbackQuery.Data, " ")
			if len(cmd) == 0 {
				continue
			}
			switch cmd[0] {
			case "remove":
				if len(cmd) != 2 {
					continue
				}
				idx, _ := strconv.Atoi(cmd[1])
				if idx < len(sh_list) {
					sh_list = append(sh_list[:idx], sh_list[idx+1:]...)
				}
				updateList()
			}
		}
	}
}

func updateList() {
	for _, v := range users {
		sendList(v)
	}
}

func sendList(chatId int64) {
	if id, ok := messages[chatId]; ok {
		bot.DeleteMessage(tgbotapi.NewDeleteMessage(chatId, id))
	}

	title := "Список покупок"
	if len(sh_list) == 0 {
		title += " пуст"
	}
	msg := tgbotapi.NewMessage(chatId, title)

	if len(sh_list) != 0 {
		var rows [][]tgbotapi.InlineKeyboardButton
		for k, v := range sh_list {
			row := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(v, "remove "+strconv.Itoa(k)))
			rows = append(rows, row)
		}
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(rows...)
	}

	if ret, err := bot.Send(msg); err == nil {
		messages[chatId] = ret.MessageID
	}
}
