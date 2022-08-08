package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dakasakti/todolist-web/config"
	"github.com/dakasakti/todolist-web/deliveries/middlewares"
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
	cookie, err := ctx.Cookie("token")
	if err != nil {
		return ctx.Redirect(http.StatusFound, "/")
	}

	user_id := middlewares.ExtractToken(cookie.Value)
	if user_id == 0 {
		return ctx.Redirect(http.StatusFound, "/login")
	}

	url := fmt.Sprintf("%s:%s/api/posts", config.GetConfig().Address, config.GetConfig().Port)
	result, err := cc.cs.GetData(url)
	if err != nil {
		return ctx.Render(http.StatusBadRequest, "index", nil)
	}

	return ctx.Render(http.StatusOK, "index", map[string]interface{}{
		"Data": result.Data,
		"Url":  fmt.Sprintf("%s:%s/logout", config.GetConfig().Address, config.GetConfig().Port),
	})
}

func (cc *clientController) Create(ctx echo.Context) error {
	cookie, err := ctx.Cookie("token")
	if err != nil {
		return ctx.Redirect(http.StatusFound, "/")
	}

	user_id := middlewares.ExtractToken(cookie.Value)
	if user_id == 0 {
		return ctx.Redirect(http.StatusFound, "/login")
	}

	return ctx.Render(http.StatusOK, "create", nil)
}

func (cc *clientController) Store(ctx echo.Context) error {
	url := fmt.Sprintf("%s:%s/api/posts", config.GetConfig().Address, config.GetConfig().Port)
	reqBody, _ := json.Marshal(map[string]interface{}{
		"description": ctx.FormValue("description"),
		"name":        ctx.FormValue("name"),
		"deadline":    ctx.FormValue("deadline"),
	})

	_, err := cc.cs.Store(url, reqBody)
	if err != nil {
		return ctx.Render(http.StatusBadRequest, "create", nil)
	}

	return ctx.Redirect(http.StatusFound, "/posts")
}

func (cc *clientController) Edit(ctx echo.Context) error {
	cookie, err := ctx.Cookie("token")
	if err != nil {
		return ctx.Redirect(http.StatusFound, "/")
	}

	user_id := middlewares.ExtractToken(cookie.Value)
	if user_id == 0 {
		return ctx.Redirect(http.StatusFound, "/login")
	}

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

func (cc *clientController) Index(ctx echo.Context) error {
	cookie, err := ctx.Cookie("token")
	if err != nil {
		return ctx.Render(http.StatusOK, "auth", nil)
	}

	user_id := middlewares.ExtractToken(cookie.Value)
	if user_id != 0 {
		return ctx.Redirect(http.StatusFound, "/posts")
	}

	return ctx.Render(http.StatusOK, "auth", nil)
}

func (cc *clientController) StoreAuth(ctx echo.Context) error {
	url := fmt.Sprintf("%s:%s/api/register", config.GetConfig().Address, config.GetConfig().Port)
	reqBody, _ := json.Marshal(map[string]interface{}{
		"fullname": ctx.FormValue("fullname"),
		"phone":    ctx.FormValue("phone"),
		"email":    ctx.FormValue("email"),
		"password": ctx.FormValue("password"),
	})

	result, err := cc.cs.Store(url, reqBody)
	fmt.Println(result)
	fmt.Println(err)
	if result.Status == 400 {
		ctx.Redirect(http.StatusFound, "/")
		return ctx.Render(http.StatusOK, "auth", result)
	}

	return ctx.Redirect(http.StatusFound, "/")
}

func (cc *clientController) LoginAuth(ctx echo.Context) error {
	url := fmt.Sprintf("%s:%s/api/login", config.GetConfig().Address, config.GetConfig().Port)
	reqBody, _ := json.Marshal(map[string]interface{}{
		"email":    ctx.FormValue("email"),
		"password": ctx.FormValue("password"),
	})

	result, err := cc.cs.Store(url, reqBody)
	fmt.Println(result)
	fmt.Println(err)
	if result.Status == 400 {
		return ctx.Render(http.StatusBadRequest, "auth", result)
	}

	ctx.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    result.Data.(string),
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	})

	return ctx.Redirect(http.StatusFound, "/posts")
}

func (cc *clientController) Logout(ctx echo.Context) error {
	ctx.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	return ctx.Redirect(http.StatusFound, "/")
}
