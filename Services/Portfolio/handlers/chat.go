package handlers

import (
	"encoding/json"
	"fmt"
	"libery_labs_portfolio/models"
	"libery_labs_portfolio/repository"
	"libery_labs_portfolio/server"
	"libery_labs_portfolio/workflows"
	"net/http"

	"github.com/Gerardo115pp/patriots_lib/echo"
)

func ChatHandler(portfolio_server server.Server) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getChatHandler(response, request)
		case http.MethodPost:
			postChatHandler(response, request)
		case http.MethodPatch:
			patchChatHandler(response, request)
		case http.MethodDelete:
			deleteChatHandler(response, request)
		case http.MethodPut:
			putChatHandler(response, request)
		case http.MethodOptions:
			response.WriteHeader(http.StatusOK)
		default:
			response.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func getChatHandler(response http.ResponseWriter, request *http.Request) {
	var claim_data *models.ChatClaims = request.Context().Value("claims").(*models.ChatClaims)

	if claim_data.ChatID == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	var chat_room *models.ChatRoom
	chat_room, err := repository.GetChatByID(claim_data.ChatID, true)
	if err != nil {
		response.WriteHeader(404)
		return
	}

	if len(chat_room.Messages) == 0 {
		chat_room.AddMessage("Good day Sir, how can I help you?", false)
	}
	fmt.Printf("Chat ID: %s\n", chat_room.ID)

	err = repository.SaveChat(chat_room)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	err = json.NewEncoder(response).Encode(chat_room)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
	}
}

func postChatHandler(response http.ResponseWriter, request *http.Request) {
	var err error
	var claim_data *models.ChatClaims = request.Context().Value("claims").(*models.ChatClaims)

	echo.Echo(echo.GreenBG, "posting to chat ID: "+claim_data.ChatID)

	if claim_data.ChatID == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	var chat_room *models.ChatRoom
	chat_room, err = repository.GetChatByID(claim_data.ChatID, false)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	var request_data map[string]string = make(map[string]string)

	err = json.NewDecoder(request.Body).Decode(&request_data)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	var message_content string = request_data["content"]

	chat_room.AddMessage(message_content, true)

	err = repository.SaveChat(chat_room)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	var new_assistant_message *models.ChatMessage
	new_assistant_message, err = workflows.SalesChatWithGPT3Turbo(chat_room) // workflows.SalesChatWithGPT3 uses AddMessage internally, don't duplicate it
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(200)

	err = json.NewEncoder(response).Encode(new_assistant_message)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
	}

	return
}
func patchChatHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}
func deleteChatHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}
func putChatHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}
