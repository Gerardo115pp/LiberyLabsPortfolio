package app_config

import (
	"fmt"
	"os"
) // Loads the configuration from the environment variables

var SERVICE_PORT string = os.Getenv("SERVICE_PORT")
var JWT_SECRET string = os.Getenv("JWT_SECRET")
var DOMAIN_SECRET string = os.Getenv("DOMAIN_SECRET")
var PROJECTS_DATA_PATH string = os.Getenv("PROJECTS_DATA_PATH")
var OPERATION_DATA_PATH string = os.Getenv("OPERATION_DATA_PATH")
var ENABLE_IDEA_GENERATION bool = os.Getenv("ENABLE_IDEA_GENERATION") == "1"

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

	PROJECTS_DATA_PATH = fmt.Sprintf("%s/%s", OPERATION_DATA_PATH, PROJECTS_DATA_PATH)
}
