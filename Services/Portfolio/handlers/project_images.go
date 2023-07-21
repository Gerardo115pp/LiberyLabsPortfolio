package handlers

import (
	"encoding/json"
	"fmt"
	"libery_labs_portfolio/repository"
	"libery_labs_portfolio/server"
	"libery_labs_portfolio/workflows"
	"net/http"

	"github.com/Gerardo115pp/patriots_lib/echo"
)

func ProjectImagesHandler(order_server server.Server) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getProjectImagesHandler(response, request)
		case http.MethodPost:
			postProjectImagesHandler(response, request)
		case http.MethodPatch:
			patchProjectImagesHandler(response, request)
		case http.MethodDelete:
			deleteProjectImagesHandler(response, request)
		case http.MethodPut:
			putProjectImagesHandler(response, request)
		case http.MethodOptions:
			response.WriteHeader(http.StatusOK)
		default:
			response.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func putProjectImagesHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func deleteProjectImagesHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func patchProjectImagesHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func postProjectImagesHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func getProjectImagesHandler(response http.ResponseWriter, request *http.Request) {

	if request.URL.Query().Get("name") != "" {
		getProjectImageByName(response, request)
		return
	} else if request.URL.Query().Get("project_id") != "" {
		getAllProjectImages(response, request)
		return
	}
}

func getProjectImageByName(response http.ResponseWriter, request *http.Request) {
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

func getAllProjectImages(response http.ResponseWriter, request *http.Request) {
	project_id := request.URL.Query().Get("project_id")
	if project_id == "" {
		echo.Echo(echo.RedFG, "Missing project_id parameter")
		response.WriteHeader(400)
		return
	}

	project_images, err := workflows.GetProjectImages(project_id)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error while getting project images: %s", err.Error()))
		response.WriteHeader(500)
		return
	}

	var response_body map[string]interface{} = make(map[string]interface{})
	response_body["mobile_count"] = len(project_images["mobile"])
	response_body["desktop_count"] = len(project_images["desktop"])
	response_body["images"] = project_images

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(response_body)
}
