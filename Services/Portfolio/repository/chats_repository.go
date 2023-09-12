package repository

import "libery_labs_portfolio/models"

type ChatRepository interface {
	GetChatByID(chat_id string) *models.ChatRoom
	SaveChat(chat *models.ChatRoom) error
}

var chat_repo_implementation ChatRepository

func SetChatRepositoryImplementation(implementation ChatRepository) {
	chat_repo_implementation = implementation
}

func GetChatByID(chat_id string) *models.ChatRoom {
	return chat_repo_implementation.GetChatByID(chat_id)
}

func SaveChat(chat *models.ChatRoom) error {
	return chat_repo_implementation.SaveChat(chat)
}
