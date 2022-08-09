package main

import (
	"fmt"

	"github.com/dakasakti/todolist-web/config"

	"github.com/dakasakti/todolist-web/deliveries/helpers"
	"github.com/dakasakti/todolist-web/deliveries/middlewares"
	"github.com/dakasakti/todolist-web/deliveries/routes"

	cm "github.com/dakasakti/todolist-web/repositories/client"
	pm "github.com/dakasakti/todolist-web/repositories/post"
	um "github.com/dakasakti/todolist-web/repositories/user"

	cs "github.com/dakasakti/todolist-web/services/client"
	ps "github.com/dakasakti/todolist-web/services/post"
	"github.com/dakasakti/todolist-web/services/renderer"
	us "github.com/dakasakti/todolist-web/services/user"
	"github.com/dakasakti/todolist-web/services/validation"

	cc "github.com/dakasakti/todolist-web/deliveries/controllers/client"
	pc "github.com/dakasakti/todolist-web/deliveries/controllers/post"
	uc "github.com/dakasakti/todolist-web/deliveries/controllers/user"

	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()
	middlewares.General(server)
	server.HTTPErrorHandler = helpers.CustomHTTPErrorHandler

	// database connection
	db := config.InitDB(*config.GetConfig())
	config.AutoMigrate(db)

	// models
	userModel := um.NewUserModel(db)
	postModel := pm.NewPostModel(db)

	// services
	validate := validation.NewValidation()
	userService := us.NewUserService(userModel)
	postService := ps.NewPostService(postModel)

	// controllers
	userController := uc.NewUserController(userService, validate)
	postController := pc.NewPostController(postService, validate)

	// routes
	routes.UserPath(server, userController)
	routes.PostPath(server, postController)

	// Register Renderer
	server.Renderer = renderer.NewRenderer()

	// Client
	clientModel := cm.NewClientModel()
	clientService := cs.NewClientService(clientModel)
	clientController := cc.NewClientController(clientService)
	routes.ClientPath(server, clientController)

	server.Logger.Fatal(server.Start(fmt.Sprintf(":%s", config.GetConfig().Port)))
}
