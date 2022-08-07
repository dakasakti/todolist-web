package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dakasakti/todolist-web/config"
	"github.com/dakasakti/todolist-web/services/client"
	"github.com/labstack/echo/v4"
)

type clientController struct {
	cs client.ClientService
}

func NewClientController(cs client.ClientService) *clientController {
	return &clientController{cs: cs}
}

func (cc *clientController) GetAll(ctx echo.Context) error {
	url := fmt.Sprintf("%s:%s/api/posts", config.GetConfig().Address, config.GetConfig().Port)

	result, err := cc.cs.GetData(url)
	if err != nil {
		return ctx.Render(http.StatusBadRequest, "index", nil)
	}

	return ctx.Render(http.StatusOK, "index", result)
}

func (cc *clientController) Create(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "create", nil)
}

func (cc *clientController) Store(ctx echo.Context) error {
	url := fmt.Sprintf("%s:%s/api/posts", config.GetConfig().Address, config.GetConfig().Port)
	reqBody, _ := json.Marshal(map[string]interface{}{
		"description": ctx.FormValue("description"),
		"name":        ctx.FormValue("name"),
		"deadline":    ctx.FormValue("deadline"),
	})

	err := cc.cs.Store(url, reqBody)
	if err != nil {
		return ctx.Render(http.StatusBadRequest, "create", nil)
	}

	return ctx.Redirect(http.StatusFound, "/posts")
}

func (cc *clientController) Edit(ctx echo.Context) error {
	url := fmt.Sprintf("%s:%s/api/posts/%s", config.GetConfig().Address, config.GetConfig().Port, ctx.Param("id"))

	result, err := cc.cs.GetData(url)
	if err != nil {
		return ctx.Render(http.StatusBadRequest, "index", nil)
	}

	return ctx.Render(http.StatusOK, "edit", result)
}

func (cc *clientController) UpdateData(ctx echo.Context) error {
	url := fmt.Sprintf("%s:%s/api/posts/%s", config.GetConfig().Address, config.GetConfig().Port, ctx.Param("id"))
	reqBody, _ := json.Marshal(map[string]interface{}{
		"description": ctx.FormValue("description"),
		"name":        ctx.FormValue("name"),
		"deadline":    ctx.FormValue("deadline"),
	})

	err := cc.cs.Update(url, reqBody)
	if err != nil {
		return ctx.Render(http.StatusBadRequest, "edit", nil)
	}

	return ctx.Redirect(http.StatusFound, "/posts")
}

func (cc *clientController) UpdateMark(ctx echo.Context) error {
	url := fmt.Sprintf("%s:%s/api/posts/%s/mark", config.GetConfig().Address, config.GetConfig().Port, ctx.Param("id"))

	err := cc.cs.Update(url, nil)
	if err != nil {
		return ctx.Render(http.StatusBadRequest, "index", nil)
	}

	return ctx.Redirect(http.StatusFound, "/posts")
}
