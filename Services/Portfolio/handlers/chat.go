package handlers

import (
	"encoding/json"
	"fmt"
	"libery_labs_portfolio/models"
	"libery_labs_portfolio/repository"
	"libery_labs_portfolio/server"
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
	echo.Echo(echo.BlueFG, fmt.Sprintf("claims: %+v", claim_data)) // copilot pendejo

	if claim_data.ChatID == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	var chat_room *models.ChatRoom = repository.GetChatByID(claim_data.ChatID)

	chat_room.AddMessage("Good day Sir, how can I help you?", false)
	fmt.Printf("Chat ID: %s\n", chat_room.ID)

	err := repository.SaveChat(chat_room)
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
	var test_message string = "{\"message\": \"Hello World\"}"
	response.WriteHeader(200)
	response.Write([]byte(test_message))
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
