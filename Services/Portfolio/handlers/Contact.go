package handlers

import (
	"encoding/json"
	"fmt"
	"libery_labs_portfolio/models"
	"libery_labs_portfolio/server"
	"libery_labs_portfolio/workflows"
	"net/http"

	"github.com/Gerardo115pp/patriots_lib/echo"
)

func ContactHandler(portfolio_server server.Server) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getContactHandler(response, request)
		case http.MethodPost:
			postContactHandler(response, request)
		case http.MethodPatch:
			patchContactHandler(response, request)
		case http.MethodDelete:
			deleteContactHandler(response, request)
		case http.MethodPut:
			putContactHandler(response, request)
		case http.MethodOptions:
			response.WriteHeader(http.StatusOK)
		default:
			response.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func getContactHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}
func postContactHandler(response http.ResponseWriter, request *http.Request) {
	var contact_message models.ContactInfo

	err := json.NewDecoder(request.Body).Decode(&contact_message)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error decoding contact message: %v", err))
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	contact_message.SenderIP = request.RemoteAddr

	err = workflows.SendTelegramMessage(contact_message.FormatTextMessage())
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error sending telegram message: %v", err))
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
}
func patchContactHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}
func deleteContactHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}
func putContactHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}
