package main

import (
	"fmt"
	"io"
	"text/template"

	"github.com/dakasakti/todolist-web/config"

	"github.com/dakasakti/todolist-web/deliveries/middlewares"
	"github.com/dakasakti/todolist-web/deliveries/routes"

	cm "github.com/dakasakti/todolist-web/repositories/client"
	pm "github.com/dakasakti/todolist-web/repositories/post"
	um "github.com/dakasakti/todolist-web/repositories/user"

	cs "github.com/dakasakti/todolist-web/services/client"
	ps "github.com/dakasakti/todolist-web/services/post"
	us "github.com/dakasakti/todolist-web/services/user"
	"github.com/dakasakti/todolist-web/services/validation"

	cc "github.com/dakasakti/todolist-web/deliveries/controllers/client"
	pc "github.com/dakasakti/todolist-web/deliveries/controllers/post"
	uc "github.com/dakasakti/todolist-web/deliveries/controllers/user"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

var funcMap = template.FuncMap{
	"itterate": func(n int) int {
		return n + 1
	},
}

func main() {
	server := echo.New()
	middlewares.General(server)

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
	renderer := &TemplateRenderer{
		templates: template.Must(template.New("views/*.html").Funcs(funcMap).ParseGlob("views/*.html")),
	}

	server.Renderer = renderer

	// Client
	clientModel := cm.NewClientModel()
	clientService := cs.NewClientService(clientModel)
	clientController := cc.NewClientController(clientService)
	routes.ClientPath(server, clientController)

	server.Logger.Fatal(server.Start(fmt.Sprintf(":%s", config.GetConfig().Port)))
}
