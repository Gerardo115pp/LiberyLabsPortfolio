package models

import "time"

type ChatMessage struct {
	Order    int    `json:"order"`
	Author   string `json:"author"`
	Content  string `json:"content"`
	SendDate string `json:"send_date"`
}

type ChatRoom struct {
	ID                 string         `json:"id"`
	CreationDate       string         `json:"creation_date"`
	LastMessageDate    string         `json:"last_message_date"`
	InstructionMessage string         `json:"instruction_message"`
	Messages           []*ChatMessage `json:"messages"`
}

func (chat *ChatRoom) AddMessage(message string, from_user bool) {
	var new_message *ChatMessage = new(ChatMessage)
	new_message.Order = len(chat.Messages)
	new_message.Content = message
	new_message.Author = "User"

	if !from_user {
		new_message.Author = "assistant"
	}

	new_message.SendDate = time.Now().Format("2006-01-02 15:04:05")
	chat.Messages = append(chat.Messages, new_message)
	chat.LastMessageDate = new_message.SendDate
}
