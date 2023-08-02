package models

import app_config "libery_labs_portfolio/Config"

type User struct {
	ID                      int64   `json:"id"`
	IsBot                   bool    `json:"is_bot"`
	FirstName               string  `json:"first_name"`
	LastName                *string `json:"last_name,omitempty"`
	Username                *string `json:"username,omitempty"`
	LanguageCode            *string `json:"language_code,omitempty"`
	IsPremium               *bool   `json:"is_premium,omitempty"`
	AddedToAttachmentMenu   *bool   `json:"added_to_attachment_menu,omitempty"`
	CanJoinGroups           *bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages *bool   `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries   *bool   `json:"supports_inline_queries,omitempty"`
}

type Chat struct {
	ID        int64   `json:"id"`
	IsBot     bool    `json:"is_bot"`
	FirstName string  `json:"first_name"`
	LastName  *string `json:"last_name,omitempty"`
	Username  *string `json:"username,omitempty"`
	Type      string  `json:"type"`
}

type Message struct {
	Id     int64  `json:"message_id"`
	Sender *User  `json:"from"`
	Chat   *Chat  `json:"chat"`
	Date   int64  `json:"date"`
	Text   string `json:"text"`
}

type NewMessage struct {
	ChatId int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func CreateNewMessage(text string) *NewMessage {
	var new_message *NewMessage = new(NewMessage)

	new_message.ChatId = app_config.TELEGRAM_CHAT_ID
	new_message.Text = text

	return new_message
}
