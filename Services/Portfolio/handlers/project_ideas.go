package handlers

import (
	"encoding/json"
	"libery_labs_portfolio/repository"
	"libery_labs_portfolio/server"
	"net/http"
	"strconv"
)

func ProjectIdeasHandler(portfolio_server server.Server) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getProjectIdeasHandler(response, request)
		case http.MethodPost:
			postProjectIdeasHandler(response, request)
		case http.MethodPatch:
			patchProjectIdeasHandler(response, request)
		case http.MethodDelete:
			deleteProjectIdeasHandler(response, request)
		case http.MethodPut:
			putProjectIdeasHandler(response, request)
		case http.MethodOptions:
			response.WriteHeader(http.StatusOK)
		default:
			response.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func postProjectIdeasHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func patchProjectIdeasHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func deleteProjectIdeasHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func putProjectIdeasHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func getProjectIdeasHandler(response http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("n") != "" {
		getNIdeasHandler(response, request)
		return
	} else {
		getAllIdeasHandler(response, request)
	}
}

func getNIdeasHandler(response http.ResponseWriter, request *http.Request) {
	var ideas_count int = 0

	ideas_count, err := strconv.Atoi(request.URL.Query().Get("n"))
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	ideas := repository.GetNIdeas(ideas_count)
	if ideas == nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	response_body := make(map[string][]string)
	response_body["project_ideas"] = ideas

	json.NewEncoder(response).Encode(response_body)
}

func getAllIdeasHandler(response http.ResponseWriter, request *http.Request) {
	ideas := repository.GetIdeas()
	if ideas == nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	response_body := make(map[string][]string)
	response_body["project_ideas"] = ideas

	json.NewEncoder(response).Encode(response_body)
}
