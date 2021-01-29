package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const (
	botToken = "1543629416:AAH3atg50vYtHA5VhgEXdIF3uFXGCsqlGSU"
	botApi   = "https://api.telegram.org/bot"
)

var (
	spisok []string
)

func main() {
	botUrl := botApi + botToken
	offset := 0
	for {
		updates, err := getUpdates(botUrl, offset)
		if err != nil {
			continue
			log.Println("Smth went wrong:", err.Error())
		}
		for _, update := range updates {
			if err = respond(botUrl, update); err != nil {
				break
			}
			fmt.Println(update.Message.Text)
			fmt.Println(update.Message.Chat.ChatId)
			fmt.Println(update.Message.From)
			spisok = append(spisok, update.Message.Text)
			offset = update.UpdateId + 1
		}
	}
}

func getUpdates(botUrl string, offset int) ([]Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var restResponse RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Result, nil
}
func respond(botUrl string, update Update) error {

	// var botMessage BotMessage
	// botMessage.ChatId = 1355658892
	// botMessage.Text = update.Message.Text
	deletemessage := Delete{update.Message.From.UserId, update.Message.MessageId}
	buf, err := json.Marshal(deletemessage)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(deletemessage)
	if _, err := http.Post(botUrl+"/deleteMessage", "application/json", bytes.NewBuffer(buf)); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
