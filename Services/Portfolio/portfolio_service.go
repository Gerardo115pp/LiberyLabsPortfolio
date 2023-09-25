package main

import (
	"context"
	"fmt"
	app_config "libery_labs_portfolio/Config"
	"libery_labs_portfolio/database"
	"libery_labs_portfolio/handlers"
	"libery_labs_portfolio/middleware"
	"libery_labs_portfolio/repository"
	"libery_labs_portfolio/server"

	"github.com/Gerardo115pp/patriot_router"
	"github.com/Gerardo115pp/patriots_lib/echo"
)

func BinderRoutes(server server.Server, router *patriot_router.Router) {
	router.RegisterRoute(patriot_router.NewRoute("/projects", true), handlers.PortfolioHandler(server))
	router.RegisterRoute(patriot_router.NewRoute("/project-images", true), handlers.ProjectImagesHandler(server))
	router.RegisterRoute(patriot_router.NewRoute("/project-ideas", true), handlers.ProjectIdeasHandler(server))
	router.RegisterRoute(patriot_router.NewRoute("/contact", true), handlers.ContactHandler(server))
	router.RegisterRoute(patriot_router.NewRoute("/is-user-human", true), handlers.GrecaptchaHandler(server))

	if app_config.SALES_CHAT_ENABLED {
		router.RegisterRoute(patriot_router.NewRoute("^/chat-claims.*", false), handlers.ChatClaimsHandler(server))
		router.RegisterRoute(patriot_router.NewRoute("/chat", true), middleware.CheckAuthCookie(handlers.ChatHandler(server)))
	}
}

func main() {
	app_config.VerifyConfig()

	echo.Echo(echo.GreenFG, "Starting libery_labs_portfolio")

	var new_server_config *server.ServerConfig = new(server.ServerConfig)
	new_server_config.Port = app_config.SERVICE_PORT

	project_repository, err := database.NewPortfolioRepository()
	if err != nil {
		echo.Echo(echo.RedFG, "Error while creating admin repository")
		echo.EchoFatal(err)
	}
	//this is usless v.1.4
	project_ideas_repository, err := database.NewProjectIdeasRepository()
	if err != nil {
		echo.Echo(echo.RedFG, "Error while creating admin repository")
		echo.EchoFatal(err)
	}

	chat_repository, err := database.NewChatDatabase()
	if err != nil {
		echo.Echo(echo.RedFG, "Error while creating chat repository")
		echo.EchoFatal(err)
	}

	repository.SetPortfolioRepositoryImplementation(project_repository)
	repository.SetProjectIdeasRepositoryImplementation(project_ideas_repository)
	repository.SetChatRepositoryImplementation(&chat_repository)

	echo.EchoDebug(fmt.Sprintf("server config: %+v", new_server_config))

	server, err := server.NewBroker(context.Background(), new_server_config)
	if err != nil {
		echo.EchoFatal(err)
	}

	server.Run(BinderRoutes)
}
