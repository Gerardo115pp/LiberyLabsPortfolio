package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	app_config "libery_labs_portfolio/Config"
	"libery_labs_portfolio/helpers"
	"libery_labs_portfolio/models"
	"os"
	"time"
)

type ChatDatabase struct {
	chats_buffer      map[string]*models.ChatRoom
	chats_buffer_size int
}

func NewChatDatabase() (ChatDatabase, error) {
	new_db := new(ChatDatabase)

	new_db.chats_buffer = make(map[string]*models.ChatRoom)
	new_db.chats_buffer_size = app_config.MAX_CHAT_SIZE

	return *new_db, nil
}

func (chat_db *ChatDatabase) GetChatByID(chat_id string) *models.ChatRoom {
	var chat_room *models.ChatRoom = new(models.ChatRoom)
	var err error
	chat_room.ID = chat_id
	chat_room.CreationDate = time.Now().Format("2006-01-02 15:04:05")
	chat_room.LastMessageDate = time.Now().Format("2006-01-02 15:04:05")
	chat_room.Messages = make([]*models.ChatMessage, 0)
	fmt.Printf("Chat ID: %s\n", chat_room.ID)

	found_chat_room := chat_db.chats_buffer[chat_id]

	if found_chat_room == nil {
		found_chat_room, err = chat_db.searchChatOnFS(chat_id)
		if err != nil {
			return nil
		}

		chat_db.chats_buffer[chat_id] = found_chat_room
	}

	fmt.Printf("Chat ID: %s\n", chat_room.ID)
	fmt.Printf("Found chat: %+v\n", found_chat_room)
	if found_chat_room != nil {
		chat_room = found_chat_room
	}
	fmt.Printf("Chat ID: %s\n", chat_room.ID)
	return chat_room
}

func (chat_db *ChatDatabase) searchChatOnFS(chat_id string) (*models.ChatRoom, error) {
	var found_chat_room *models.ChatRoom

	// If chat exists on the file system
	if helpers.FileExists(fmt.Sprintf("%s/%s", app_config.CHATS_DATA_PATH, chat_id)) {
		found_chat_room = new(models.ChatRoom)
		chat_data, err := ioutil.ReadFile(fmt.Sprintf("%s/%s/chat.json", app_config.CHATS_DATA_PATH, chat_id))
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(chat_data, found_chat_room)
		if err != nil {
			return nil, err
		}
	}

	return found_chat_room, nil
}

func (chat_db *ChatDatabase) SaveChat(chat *models.ChatRoom) error {
	chat_data, err := json.Marshal(chat)
	if err != nil {
		return err
	}

	chat_directory := fmt.Sprintf("%s/%s", app_config.CHATS_DATA_PATH, chat.ID)

	if !helpers.FileExists(chat_directory) {
		err = os.Mkdir(chat_directory, os.ModePerm) // 0777
		if err != nil {
			return err
		}
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s/chat.json", chat_directory), chat_data, 0777)
	if err != nil {
		return err
	}

	return nil
}
