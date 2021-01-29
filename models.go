package main

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Chat      Chat   `json:"chat"`
	Text      string `json:"text"`
	From      User   `json:"from"`
	MessageId int    `json:"message_id"`
}

type Delete struct {
	ChatId    int `json:"chat_id"`
	MessageId int `json:"message_id"`
}

type User struct {
	UserId int    `json:"id"`
	Name   string `json:"username"`
}

type Chat struct {
	ChatId int `json:"chat"`
}

type RestResponse struct {
	Result []Update `json:"result"`
}

type BotMessage struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
}
