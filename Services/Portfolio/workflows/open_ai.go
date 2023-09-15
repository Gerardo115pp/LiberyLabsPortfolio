package workflows

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"libery_labs_portfolio/models"
	"libery_labs_portfolio/repository"
	"net/http"
	"os"

	app_config "libery_labs_portfolio/Config"

	"github.com/Gerardo115pp/patriots_lib/echo"
)

func CreateNewProjectIdea() (string, error) {
	idea_examples := repository.GetNIdeas(4)

	var prompt string = "Give me a single full stack project idea, the idea should be a brief sentence that describes the project in a compelling way. you response should absolutely not include anything aside from the idea sentence itself and this sentence must be shorter than ~40 characters. This is a list of four example responses: "
	for _, idea := range idea_examples {
		prompt += idea
		prompt += ", "
	}

	project_idea, err := requestGPT3TurboIdea(prompt)

	return project_idea, err
}

func SalesChatWithGPT3Turbo(chat_room *models.ChatRoom) (*models.ChatMessage, error) {
	var chat_request *models.GPT3TurboRequest = models.CreateGPT3TurboRequest()

	chat_request.Temperature = app_config.SALES_CHAT_TEMPERATURE
	chat_request.TopP = app_config.SALES_CHAT_TOP_P
	chat_request.MaxTokens = app_config.SALES_CHAT_MAX_TOKENS

	var system_message *models.GPT3TurboMessage = new(models.GPT3TurboMessage)

	system_message.Role = "system"
	system_message.Content = app_config.SALES_CHAT_INSTRUCTION

	chat_request.Messages = append(chat_request.Messages, system_message)

	var new_message *models.GPT3TurboMessage

	for _, message := range chat_room.Messages {
		new_message = new(models.GPT3TurboMessage)
		new_message.Role = message.Author
		new_message.Content = message.Content

		chat_request.Messages = append(chat_request.Messages, new_message)
	}

	response, err := requestGPT3Turbo(chat_request)
	if err != nil {
		return nil, err
	}

	new_chat_message := chat_room.AddMessage(response.Choices[0].Message.Content, false)

	return new_chat_message, nil
}

func requestGPT3Turbo(request *models.GPT3TurboRequest) (*models.GPT3TurboResponse, error) {
	var request_body []byte
	var err error

	request_body, err = json.Marshal(request)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error marshalling request body: %s", err.Error()))
		return nil, err
	}

	body := bytes.NewBuffer(request_body)
	http_request, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", body)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error sending GPT 3-Turbo request: %s", err.Error()))
		return nil, err
	}

	http_request.Header.Add("Content-Type", "application/json")
	http_request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", app_config.OPENAI_API_KEY))

	http_client := http.Client{}
	response, err := http_client.Do(http_request)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error sending GPT 3-Turbo request: %s", err.Error()))
		return nil, err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error sending GPT 3-Turbo request, status code: %d", response.StatusCode))
		body_content, err := ioutil.ReadAll(response.Body)
		if err != nil {
			echo.Echo(echo.RedFG, fmt.Sprintf("Error reading response body: %s", err.Error()))
			return nil, err
		}
		echo.Echo(echo.RedFG, fmt.Sprintf("Reason: %s", string(body_content)))
		echo.Echo(echo.RedFG, fmt.Sprintf("Request body: %s", string(request_body)))
		return nil, fmt.Errorf("Error sending GPT 3-Turbo request, status code: %d", response.StatusCode)
	}

	return parseGPT3TurboResponse(response)
}

func requestGPT3TurboIdea(prompt string) (string, error) {
	request_data := map[string]interface{}{
		"max_tokens":  40,
		"temperature": 1.26,
		"model":       "gpt-3.5-turbo",
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": prompt,
			},
		},
	}

	request_body, err := json.Marshal(request_data)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error marshalling request body: %s", err.Error()))
		return "", err
	}

	body := bytes.NewBuffer(request_body)
	request, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", body)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error creating new project idea: %s", err.Error()))
		return "", err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("OPENAI_API_KEY")))

	http_client := http.Client{}
	response, err := http_client.Do(request)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error creating new project idea: %s", err.Error()))
		return "", err
	}

	if response.StatusCode != 200 {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error creating new project idea, status code: %d", response.StatusCode))
		body_content, err := ioutil.ReadAll(response.Body)
		if err != nil {
			echo.Echo(echo.RedFG, fmt.Sprintf("Error reading response body: %s", err.Error()))
			return "", err
		}
		echo.Echo(echo.RedFG, fmt.Sprintf("Reason: %s", string(body_content)))
		echo.Echo(echo.RedFG, fmt.Sprintf("Request body: %s", string(request_body)))
		return "", err
	}

	response_data, err := parseGPT3TurboResponse(response)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error parsing response: %s", err.Error()))
		return "", err
	}

	return response_data.Choices[0].Message.Content, nil
}

func parseGPT3TurboResponse(response *http.Response) (*models.GPT3TurboResponse, error) {
	var response_data *models.GPT3TurboResponse = new(models.GPT3TurboResponse)
	err := json.NewDecoder(response.Body).Decode(&response_data)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error decoding response: %s", err.Error()))
		return nil, err
	}

	return response_data, nil
}
