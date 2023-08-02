package workflows

import (
	"bytes"
	"encoding/json"
	"fmt"
	app_config "libery_labs_portfolio/Config"
	"libery_labs_portfolio/models"
	"net/http"
)

func TestBot() error {
	bot_token := fmt.Sprintf("bot%s", app_config.TELEGRAM_API_KEY)
	http_client := http.Client{}

	request, err := http.NewRequest("GET", fmt.Sprintf("https://api.telegram.org/%s/getMe", bot_token), nil)
	if err != nil {
		return err
	}

	response, err := http_client.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return fmt.Errorf("Error: %s", response.Status)
	}

	response_body := &struct {
		Ok     bool         `json:"ok"`
		Result *models.User `json:"result"`
	}{}

	err = json.NewDecoder(response.Body).Decode(&response_body)
	if err != nil {
		return err
	}

	if !response_body.Ok {
		return fmt.Errorf("Error: %s", response.Status)
	}

	fmt.Printf("Bot username: %s\n", *response_body.Result.Username)

	return nil
}

func SendTelegramMessage(text string) error {
	var err error
	var message *models.NewMessage = models.CreateNewMessage(text)

	bot_token := fmt.Sprintf("bot%s", app_config.TELEGRAM_API_KEY)

	http_client := http.Client{}

	request_body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(request_body)

	request, err := http.NewRequest("POST", fmt.Sprintf("https://api.telegram.org/%s/sendMessage", bot_token), body)
	if err != nil {
		return err
	}

	request.Header.Add("Content-Type", "application/json")

	response, err := http_client.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return fmt.Errorf("Error: %s", response.Status)
	}

	return nil
}
