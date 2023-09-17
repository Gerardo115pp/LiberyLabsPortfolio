package handlers

import (
	"encoding/json"
	"libery_labs_portfolio/helpers"
	"libery_labs_portfolio/server"
	"libery_labs_portfolio/workflows"
	"net/http"

	"github.com/Gerardo115pp/patriots_lib/echo"
)

func GrecaptchaHandler(portfolio_server server.Server) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		var current_origin string = request.Host

		echo.Echo(echo.GreenFG, "Request from origin: "+current_origin)

		if current_origin != "libery-labs.com" && current_origin != "developer-libery-labs.com" {
			response.WriteHeader(http.StatusForbidden)
			return
		}

		response.Header().Set("Access-Control-Allow-Origin", current_origin)

		switch request.Method {
		case http.MethodGet:
			getGrecaptchaHandler(response, request)
		case http.MethodPost:
			postGrecaptchaHandler(response, request)
		case http.MethodPatch:
			patchGrecaptchaHandler(response, request)
		case http.MethodDelete:
			deleteGrecaptchaHandler(response, request)
		case http.MethodPut:
			putGrecaptchaHandler(response, request)
		case http.MethodOptions:
			response.WriteHeader(http.StatusOK)
		default:
			response.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func getGrecaptchaHandler(response http.ResponseWriter, request *http.Request) {
	var grecaptcha_user_response string = request.URL.Query().Get("token")

	if grecaptcha_user_response == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	is_human, err := workflows.Google.VerifyCaptcha(grecaptcha_user_response)
	if err != nil {
		helpers.WriteRejection(response, 400, "bad-user-response")
		return
	}

	is_human_response := &struct{ IsHuman bool }{IsHuman: is_human}

	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(is_human_response)

	return
}
func postGrecaptchaHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}
func patchGrecaptchaHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}
func deleteGrecaptchaHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}
func putGrecaptchaHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}
