package handlers

import (
	"libery_labs_portfolio/server"
	"net/http"
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
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}
func postChatHandler(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("{\"message\": \"Hello World\"}"))
	response.WriteHeader(200)
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
