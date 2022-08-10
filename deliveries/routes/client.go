package routes

import (
	"github.com/dakasakti/todolist-web/deliveries/controllers/client"
	"github.com/labstack/echo/v4"
)

func ClientPath(e *echo.Echo, cc client.ClientController) {
	// middlewares.JWTSign()
	e.GET("/posts", cc.GetAll)
	e.GET("/posts/create", cc.Create)
	e.POST("/posts", cc.Store)
	e.GET("/posts/:id/edit", cc.Edit)
	e.POST("posts/:id", cc.UpdateData)
	e.POST("/posts/:id/mark", cc.UpdateMark)

	// Auth
	e.GET("/", cc.Index)
	e.GET("/register", cc.Index)
	e.POST("/register", cc.Register)
	e.GET("/login", cc.Index)
	e.POST("/login", cc.Login)
	e.POST("/logout", cc.Logout)
}
