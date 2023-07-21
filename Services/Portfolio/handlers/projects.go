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

func PortfolioHandler(order_server server.Server) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getPortfolioHandler(response, request)
		case http.MethodPost:
			postPortfolioHandler(response, request)
		case http.MethodPatch:
			patchPortfolioHandler(response, request)
		case http.MethodDelete:
			deletePortfolioHandler(response, request)
		case http.MethodPut:
			putPortfolioHandler(response, request)
		case http.MethodOptions:
			response.WriteHeader(http.StatusOK)
		default:
			response.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func putPortfolioHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func deletePortfolioHandler(response http.ResponseWriter, request *http.Request) {
	var request_body map[string]string = make(map[string]string)

	err := json.NewDecoder(request.Body).Decode(&request_body)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error decoding delete Portfolio request: %s", err.Error()))
		response.WriteHeader(400)
		return
	}

	project_id, exists := request_body["id"]
	if !exists {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error decoding delete Portfolio request: %s", err.Error()))
		response.WriteHeader(400)
		return
	}

	err = repository.DeleteProject(request.Context(), project_id)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error deleting Portfolio: %s", err.Error()))
		response.WriteHeader(500)
		return
	}

	response.WriteHeader(200)
}

func patchPortfolioHandler(response http.ResponseWriter, request *http.Request) {
	var update_project_request map[string]any
	err := json.NewDecoder(request.Body).Decode(&update_project_request)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error decoding update Portfolio request: %s", err.Error()))
		response.WriteHeader(400)
		return
	}

	target_project, err := repository.GetProjectByID(request.Context(), update_project_request["id"].(string))
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error getting Portfolio by id: %s", err.Error()))
		response.WriteHeader(400)
		return
	}

	target_project.Update(update_project_request)

	err = repository.UpdateProject(request.Context(), target_project)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error updating Portfolio: %s", err.Error()))
		response.WriteHeader(400)
		return
	}

	response.WriteHeader(200)
}

func postPortfolioHandler(response http.ResponseWriter, request *http.Request) {

	new_project := new(models.Project)

	err := json.NewDecoder(request.Body).Decode(&new_project)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error decoding new Portfolio request: %s", err.Error()))
		response.WriteHeader(400)
		return
	}

	repository.InsertProject(request.Context(), *new_project)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(201)
}

func getPortfolioHandler(response http.ResponseWriter, request *http.Request) {

	if request.URL.Query().Get("id") != "" {
		getPortfolioByIDHandler(response, request)
		return
	} else {
		getAllPortfoliosHandler(response, request)
		return
	}
}

func getPortfolioByIDHandler(response http.ResponseWriter, request *http.Request) {
	var portfolio_id string = request.URL.Query().Get("id")
	if portfolio_id == "" {
		response.WriteHeader(400)
		return
	}

	portfolio, err := repository.GetProjectByID(request.Context(), portfolio_id)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error while getting Portfolio by id: %s", err.Error()))
		response.WriteHeader(404)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(portfolio)
	response.WriteHeader(200)
}

func getAllPortfoliosHandler(response http.ResponseWriter, request *http.Request) {
	portfolios, err := repository.GetProjects(request.Context())
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error while getting Portfolios: %s", err.Error()))
		response.WriteHeader(404)
		return
	}

	projects_clean_data := make([]map[string]any, len(portfolios))
	for i, project := range portfolios {
		clean_data := make(map[string]any)
		clean_data["id"] = project.ID
		clean_data["name"] = project.Name
		clean_data["tech_stack"] = project.TechStack
		clean_data["relevance"] = project.Relevance
		clean_data["type"] = project.Type

		projects_clean_data[i] = clean_data
	}

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(projects_clean_data)
	response.WriteHeader(200)
}
