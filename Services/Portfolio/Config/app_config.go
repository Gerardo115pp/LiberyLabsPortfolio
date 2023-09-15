package app_config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/Gerardo115pp/patriots_lib/echo"
) // Loads the configuration from the environment variables

var SERVICE_PORT string = os.Getenv("SERVICE_PORT")
var JWT_SECRET string = os.Getenv("JWT_SECRET")
var DOMAIN_SECRET string = os.Getenv("DOMAIN_SECRET")
var PROJECTS_DATA_PATH string = os.Getenv("PROJECTS_DATA_PATH")
var OPERATION_DATA_PATH string = os.Getenv("OPERATION_DATA_PATH")
var OPENAI_API_KEY string = os.Getenv("OPENAI_API_KEY")
var ENABLE_IDEA_GENERATION bool = os.Getenv("ENABLE_IDEA_GENERATION") == "1"
var TELEGRAM_API_KEY string = os.Getenv("TELEGRAM_API_KEY")
var chat_id string = os.Getenv("TELEGRAM_CHAT_ID")
var TELEGRAM_CHAT_ID int64 = 0

// -------- Asisstant chats config --------
var CHAT_CLAIM_COOKIE_NAME string = "portfolio_libery_chat_token"
var MAX_CHAT_SIZE int = 10
var custom_max_chat_size string = os.Getenv("MAX_CHAT_SIZE")
var CHATS_DATA_PATH string = fmt.Sprintf("%s/chats", OPERATION_DATA_PATH)

// ------- Settings config -------
var SALES_CHAT_INSTRUCTION string
var SALES_CHAT_ENABLED bool
var SALES_CHAT_TEMPERATURE float64 = 1.2
var SALES_CHAT_TOP_P float64 = 1.0
var SALES_CHAT_MAX_TOKENS int = 100

func VerifyConfig() {

	if SERVICE_PORT == "" {
		panic("CUST_PORT environment variable is required")
	}
	if JWT_SECRET == "" {
		panic("JWT_SECRET environment variable is required")
	}
	if DOMAIN_SECRET == "" {
		panic("DOMAIN_SECRET environment variable is required")
	}
	if PROJECTS_DATA_PATH == "" {
		PROJECTS_DATA_PATH = "projects"
	}
	if OPERATION_DATA_PATH == "" {
		panic("OPERATION_DATA_PATH environment variable is required")
	}
	if OPENAI_API_KEY == "" {
		panic("OPENAI_API_KEY environment variable is required")
	}
	if TELEGRAM_API_KEY == "" {
		panic("TELEGRAM_API_KEY environment variable is required")
	}
	if chat_id == "" {
		panic("TELEGRAM_CHAT_ID environment variable is required")
	}
	var err error

	TELEGRAM_CHAT_ID, err = strconv.ParseInt(chat_id, 10, 64)
	if err != nil {
		panic(err)
	}

	if custom_max_chat_size != "" {
		MAX_CHAT_SIZE, err = strconv.Atoi(custom_max_chat_size)
		if err != nil {
			echo.EchoFatal(err)
		}
	}

	PROJECTS_DATA_PATH = fmt.Sprintf("%s/%s", OPERATION_DATA_PATH, PROJECTS_DATA_PATH)

	err = loadSettings()
	if err != nil {
		echo.Echo(echo.RedFG, "Error loading settings: %s", err.Error())
		echo.EchoFatal(err)
	}
}

func loadSettings() error {
	var settings_path string = fmt.Sprintf("%s/settings.json", OPERATION_DATA_PATH)
	var err error

	if _, err = os.Stat(settings_path); os.IsNotExist(err) {
		return err
	}

	var settings_content []byte

	settings_content, err = ioutil.ReadFile(settings_path)
	if err != nil {
		return err
	}

	var settings map[string]interface{}

	err = json.Unmarshal(settings_content, &settings)
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			echo.EchoFatal(fmt.Errorf("Error loading settings: %s", r))
		}
	}()

	var sales_chat_settings map[string]interface{} = settings["sales_chat"].(map[string]interface{})

	SALES_CHAT_INSTRUCTION = sales_chat_settings["instruction_message"].(string)
	if SALES_CHAT_INSTRUCTION == "" {
		return fmt.Errorf("sales_chat_instruction is required")
	}

	if sales_chat_settings["chat_enabled"] == nil {
		return fmt.Errorf("chat_enabled is required")
	}
	SALES_CHAT_ENABLED = sales_chat_settings["chat_enabled"].(bool)

	if sales_chat_settings["temperature"] != nil {
		SALES_CHAT_TEMPERATURE = sales_chat_settings["temperature"].(float64)
	}

	if sales_chat_settings["top_p"] != nil {
		SALES_CHAT_TOP_P = sales_chat_settings["top_p"].(float64)
	}

	if sales_chat_settings["max_tokens"] != nil {
		SALES_CHAT_MAX_TOKENS = int(sales_chat_settings["max_tokens"].(float64))
	}

	return nil
}
