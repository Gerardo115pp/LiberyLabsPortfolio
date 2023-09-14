package handlers

import (
	"encoding/json"
	"libery_labs_portfolio/repository"
	"libery_labs_portfolio/server"
	"libery_labs_portfolio/workflows"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/Gerardo115pp/patriots_lib/echo"
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
	} else if request.URL.Query().Get("mode") == "r" {
		getRandomIdeasHandler(response, request)
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

func getRandomIdeasHandler(response http.ResponseWriter, request *http.Request) {
	var ideas_count int = repository.IdeasCount()
	var new_idea_chance float64 = 1.0 - (float64(ideas_count) / 100.0)
	var max_chance float64 = 0.65

	if new_idea_chance > max_chance {
		new_idea_chance = max_chance
	}

	var idea string
	var err error
	rand.Seed(time.Now().UnixNano())

	random_roll := rand.Float64()

	echo.Echo(echo.YellowBG, "Random roll: "+strconv.FormatFloat(random_roll, 'f', 6, 64)+"Chance: "+strconv.FormatFloat(new_idea_chance, 'f', 6, 64))

	if random_roll <= new_idea_chance {
		echo.Echo(echo.GreenBG, "New project idea generated")
		idea, err = workflows.CreateNewProjectIdea()
		if err != nil {
			echo.EchoErr(err)
			response.WriteHeader(http.StatusInternalServerError)
			return
		}

		repository.AddIdea(idea)
	} else {
		idea = repository.GetRandomIdea()
	}

	if idea == "" {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	var response_body map[string]string = make(map[string]string)
	response_body["project_idea"] = idea

	json.NewEncoder(response).Encode(response_body)
}
