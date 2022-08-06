package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"text/template"

	"github.com/dakasakti/postingan/config"

	"github.com/dakasakti/postingan/deliveries/helpers"
	"github.com/dakasakti/postingan/deliveries/middlewares"
	"github.com/dakasakti/postingan/deliveries/routes"

	pm "github.com/dakasakti/postingan/repositories/post"
	um "github.com/dakasakti/postingan/repositories/user"

	ps "github.com/dakasakti/postingan/services/post"
	us "github.com/dakasakti/postingan/services/user"
	"github.com/dakasakti/postingan/services/validation"

	pc "github.com/dakasakti/postingan/deliveries/controllers/post"
	uc "github.com/dakasakti/postingan/deliveries/controllers/user"

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

	renderer := &TemplateRenderer{
		templates: template.Must(template.New("views/*.html").Funcs(funcMap).ParseGlob("views/*.html")),
	}
	server.Renderer = renderer

	server.GET("/posts/create", func(c echo.Context) error {
		return c.Render(http.StatusOK, "create", nil)
	})

	server.POST("/posts", func(c echo.Context) error {
		requestBody, _ := json.Marshal(map[string]interface{}{
			"description": c.FormValue("description"),
			"name":        c.FormValue("name"),
			"deadline":    c.FormValue("deadline"),
		})

		url := fmt.Sprintf("%s:%s/api/posts", config.GetConfig().Address, config.GetConfig().Port)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		if err != nil {
			fmt.Println(err.Error())
			return c.Render(http.StatusBadRequest, "create", nil)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err.Error())
			return c.Render(http.StatusBadRequest, "create", nil)
		}

		defer resp.Body.Close()

		var result helpers.ResponseJSON
		json.NewDecoder(resp.Body).Decode(&result)
		return c.Redirect(http.StatusFound, "/posts")
	})

	server.GET("/posts", func(c echo.Context) error {
		url := fmt.Sprintf("%s:%s/api/posts", config.GetConfig().Address, config.GetConfig().Port)
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		if err != nil {
			fmt.Println(err.Error())
			return c.Render(http.StatusBadRequest, "index", nil)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err.Error())
			return c.Render(http.StatusBadRequest, "index", nil)
		}

		defer resp.Body.Close()

		var result helpers.ResponseJSON
		json.NewDecoder(resp.Body).Decode(&result)
		return c.Render(http.StatusOK, "index", result)
	})

	server.POST("/posts/:id/mark", func(c echo.Context) error {
		id := c.Param("id")
		url := fmt.Sprintf("%s:%s/api/posts/%s/mark", config.GetConfig().Address, config.GetConfig().Port, id)
		req, err := http.NewRequest("PUT", url, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		if err != nil {
			fmt.Println(err.Error())
			return c.Render(http.StatusBadRequest, "index", nil)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err.Error())
			return c.Render(http.StatusBadRequest, "index", nil)
		}

		defer resp.Body.Close()

		var result helpers.ResponseJSON
		json.NewDecoder(resp.Body).Decode(&result)
		return c.Redirect(http.StatusFound, "/posts")
	})

	server.GET("/posts/:id/edit", func(c echo.Context) error {
		id := c.Param("id")
		url := fmt.Sprintf("%s:%s/api/posts/%s", config.GetConfig().Address, config.GetConfig().Port, id)
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		if err != nil {
			fmt.Println(err.Error())
			return c.Render(http.StatusBadRequest, "index", nil)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err.Error())
			return c.Render(http.StatusBadRequest, "index", nil)
		}

		defer resp.Body.Close()

		var result helpers.ResponseJSON
		json.NewDecoder(resp.Body).Decode(&result)
		return c.Render(http.StatusOK, "edit", result)
	})

	server.POST("/posts/:id", func(c echo.Context) error {
		id := c.Param("id")
		url := fmt.Sprintf("%s:%s/api/posts/%s", config.GetConfig().Address, config.GetConfig().Port, id)

		requestBody, _ := json.Marshal(map[string]interface{}{
			"description": c.FormValue("description"),
			"name":        c.FormValue("name"),
			"deadline":    c.FormValue("deadline"),
		})

		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		if err != nil {
			fmt.Println(err.Error())
			return c.Render(http.StatusBadRequest, "index", nil)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err.Error())
			return c.Render(http.StatusBadRequest, "index", nil)
		}

		defer resp.Body.Close()

		var result helpers.ResponseJSON
		json.NewDecoder(resp.Body).Decode(&result)
		return c.Redirect(http.StatusFound, "/posts")
	})

	server.Logger.Fatal(server.Start(fmt.Sprintf(":%s", config.GetConfig().Port)))
}
