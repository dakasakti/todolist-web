package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dakasakti/todolist-web/config"
	"github.com/dakasakti/todolist-web/deliveries/middlewares"
	"github.com/dakasakti/todolist-web/services/client"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
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
	if result.Status == 400 || err != nil {
		fmt.Println(result.Message)
		return ctx.Render(http.StatusBadRequest, "index", nil)
	}

	url = fmt.Sprintf("%s:%s/api/profile", config.GetConfig().Address, config.GetConfig().Port)
	dataUser, err := cc.cs.GetDatawithAuth(url, cookie.Value)
	if dataUser.Status == 400 || err != nil {
		return ctx.Render(http.StatusBadRequest, "index", nil)
	}

	sess, _ := session.Get("session", ctx)
	return ctx.Render(http.StatusOK, "index", map[string]interface{}{
		"Data":    result.Data,
		"User":    dataUser.Data,
		"Message": sess.Values["message"],
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
	cookie, err := ctx.Cookie("token")
	if err != nil {
		return ctx.Redirect(http.StatusFound, "/")
	}

	url := fmt.Sprintf("%s:%s/api/posts", config.GetConfig().Address, config.GetConfig().Port)
	reqBody, _ := json.Marshal(map[string]interface{}{
		"description": ctx.FormValue("description"),
		"name":        ctx.FormValue("name"),
		"deadline":    ctx.FormValue("deadline"),
	})

	result, err := cc.cs.StorewithAuth(url, cookie.Value, reqBody)
	if result.Status == 400 || err != nil {
		fmt.Println(result.Message)
		return ctx.Render(http.StatusOK, "create", map[string]interface{}{
			"Data": result.Data,
		})
	}

	sess, _ := session.Get("session", ctx)
	sess.Options = &sessions.Options{
		Path:   "/posts",
		MaxAge: 5,
	}

	sess.Values["message"] = result.Message
	sess.Save(ctx.Request(), ctx.Response())
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
	if result.Status == 400 || err != nil {
		fmt.Println(result.Message)
		return ctx.Render(http.StatusBadRequest, "index", nil)
	}

	return ctx.Render(http.StatusOK, "edit", map[string]interface{}{
		"Data": result.Data,
	})
}

func (cc *clientController) UpdateData(ctx echo.Context) error {
	cookie, err := ctx.Cookie("token")
	if err != nil {
		return ctx.Redirect(http.StatusFound, "/")
	}

	url := fmt.Sprintf("%s:%s/api/posts/%s", config.GetConfig().Address, config.GetConfig().Port, ctx.Param("id"))
	reqBody, _ := json.Marshal(map[string]interface{}{
		"description": ctx.FormValue("description"),
		"name":        ctx.FormValue("name"),
		"deadline":    ctx.FormValue("deadline"),
	})

	result, err := cc.cs.UpdatewithAuth(url, cookie.Value, reqBody)
	if result.Status == 400 || err != nil {
		fmt.Println(result.Message)
		return ctx.Render(http.StatusOK, "edit", map[string]interface{}{
			"Data": result.Data,
		})
	}

	sess, _ := session.Get("session", ctx)
	sess.Options = &sessions.Options{
		Path:   "/posts",
		MaxAge: 5,
	}

	sess.Values["message"] = result.Message
	sess.Save(ctx.Request(), ctx.Response())
	return ctx.Redirect(http.StatusFound, "/posts")
}

func (cc *clientController) UpdateMark(ctx echo.Context) error {
	cookie, err := ctx.Cookie("token")
	if err != nil {
		return ctx.Redirect(http.StatusFound, "/")
	}

	url := fmt.Sprintf("%s:%s/api/posts/%s/mark", config.GetConfig().Address, config.GetConfig().Port, ctx.Param("id"))
	result, err := cc.cs.UpdatewithAuth(url, cookie.Value, nil)
	if result.Status == 400 || err != nil {
		fmt.Println(result.Message)
		return ctx.Render(http.StatusOK, "index", result)
	}

	sess, _ := session.Get("session", ctx)
	sess.Options = &sessions.Options{
		Path:   "/posts",
		MaxAge: 5,
	}

	sess.Values["message"] = result.Message
	sess.Save(ctx.Request(), ctx.Response())
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

	sess, _ := session.Get("session", ctx)
	return ctx.Render(http.StatusOK, "auth", map[string]interface{}{
		"Message": sess.Values["message"],
	})
}

func (cc *clientController) Register(ctx echo.Context) error {
	url := fmt.Sprintf("%s:%s/api/register", config.GetConfig().Address, config.GetConfig().Port)
	reqBody, _ := json.Marshal(map[string]interface{}{
		"fullname": ctx.FormValue("fullname"),
		"phone":    ctx.FormValue("phone"),
		"email":    ctx.FormValue("email"),
		"password": ctx.FormValue("password"),
	})

	result, err := cc.cs.Store(url, reqBody)
	if result.Status == 400 || err != nil {
		fmt.Println(result.Message)
		return ctx.Render(http.StatusOK, "auth", map[string]interface{}{
			"ErrorData": result.Data,
		})
	}

	sess, _ := session.Get("session", ctx)
	sess.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 10,
	}

	sess.Values["message"] = result.Message
	sess.Save(ctx.Request(), ctx.Response())
	return ctx.Redirect(http.StatusFound, "/")
}

func (cc *clientController) Login(ctx echo.Context) error {
	url := fmt.Sprintf("%s:%s/api/login", config.GetConfig().Address, config.GetConfig().Port)
	reqBody, _ := json.Marshal(map[string]interface{}{
		"email":    ctx.FormValue("email"),
		"password": ctx.FormValue("password"),
	})

	result, err := cc.cs.Store(url, reqBody)
	if result.Status == 400 || err != nil {
		fmt.Println(result.Message)
		return ctx.Render(http.StatusOK, "auth", map[string]interface{}{
			"Data": result.Data,
		})
	}

	if result.Status == 401 || err != nil {
		fmt.Println(result.Message)
		return ctx.Render(http.StatusOK, "auth", map[string]interface{}{
			"Message": result.Message,
		})
	}

	sess, _ := session.Get("session", ctx)
	sess.Options = &sessions.Options{
		Path:   "/posts",
		MaxAge: 5,
	}

	sess.Values["message"] = result.Message
	sess.Save(ctx.Request(), ctx.Response())

	ctx.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   result.Data.(string),
		Expires: time.Now().Add(time.Minute * 5),
	})

	return ctx.Redirect(http.StatusFound, "/posts")
}

func (cc *clientController) Logout(ctx echo.Context) error {
	ctx.SetCookie(&http.Cookie{
		Name:   "token",
		Value:  "",
		MaxAge: -1,
	})

	return ctx.Redirect(http.StatusFound, "/")
}
