package main

import (
	"context"
	"fmt"
	app_config "libery_labs_portfolio/Config"
	"libery_labs_portfolio/database"
	"libery_labs_portfolio/handlers"
	"libery_labs_portfolio/repository"
	"libery_labs_portfolio/server"

	"github.com/Gerardo115pp/patriot_router"
	"github.com/Gerardo115pp/patriots_lib/echo"
)

func BinderRoutes(server server.Server, router *patriot_router.Router) {
	router.RegisterRoute(patriot_router.NewRoute("/projects", true), handlers.PortfolioHandler(server))
	router.RegisterRoute(patriot_router.NewRoute("/project-images", true), handlers.ProjectImagesHandler(server))
	// router.RegisterRoute(patriot_router.NewRoute("^/profile_pictures.*", false), middleware.CheckAuth(handlers.ProfilePicturesHandler(server)))
}

func main() {
	app_config.VerifyConfig()

	echo.Echo(echo.GreenFG, "Starting libery_labs_portfolio")

	var new_server_config *server.ServerConfig = new(server.ServerConfig)
	new_server_config.Port = app_config.SERVICE_PORT

	admin_repository, err := database.NewPortfolioRepository()
	if err != nil {
		echo.Echo(echo.RedFG, "Error while creating admin repository")
		echo.EchoFatal(err)
	}

	repository.SetPortfolioRepositoryImplementation(admin_repository)

	echo.EchoDebug(fmt.Sprintf("server config: %+v", new_server_config))

	server, err := server.NewBroker(context.Background(), new_server_config)
	if err != nil {
		echo.EchoFatal(err)
	}

	server.Run(BinderRoutes)
}
