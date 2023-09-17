package workflows

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	app_config "libery_labs_portfolio/Config"
	"libery_labs_portfolio/models"
	"net/http"

	"github.com/Gerardo115pp/patriots_lib/echo"
)

type GoogleWorkflow struct {
	recaptcha_sk string
}

var Google *GoogleWorkflow = &GoogleWorkflow{recaptcha_sk: app_config.RECAPTCHA_SK}

func (google_ws GoogleWorkflow) VerifyCaptcha(token string) (bool, error) {

	var verify_params string = fmt.Sprintf("secret=%s&response=%s", google_ws.recaptcha_sk, token)

	http_request, err := http.NewRequest("POST", fmt.Sprintf("https://www.google.com/recaptcha/api/siteverify?%s", verify_params), nil)
	if err != nil {
		return false, err
	}

	http_request.Header.Add("Content-Type", "application/json")

	http_client := http.Client{}
	response, err := http_client.Do(http_request)
	if err != nil {
		return false, err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return false, err
	}

	verify_response, err := parseVerifyCaptchaResponse(response)
	if err != nil {
		return false, err
	}

	if !verify_response.Success {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error verifying captcha: %+v", verify_response.ErrorCodes))
	}

	return verify_response.Success, nil
}

func parseVerifyCaptchaResponse(response *http.Response) (*models.RecaptchaVerifyResponse, error) {
	var verify_response *models.RecaptchaVerifyResponse = new(models.RecaptchaVerifyResponse)
	var err error

	body_content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body_content, verify_response)
	if err != nil {
		return nil, err
	}

	return verify_response, nil
}
