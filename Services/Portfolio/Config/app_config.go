package app_config

import (
	"fmt"
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
}
