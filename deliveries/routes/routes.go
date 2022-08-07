package routes

import (
	"github.com/labstack/echo/v4"

	ps "github.com/dakasakti/todolist-web/deliveries/controllers/post"
	uc "github.com/dakasakti/todolist-web/deliveries/controllers/user"
)

func UserPath(e *echo.Echo, uc uc.UserController) {
	api := e.Group("/api")
	api.POST("/register", uc.Register)
	api.POST("/login", uc.Login)
}

func PostPath(e *echo.Echo, ps ps.PostController) {
	// middlewares.JWTSign()
	api := e.Group("/api")
	api.POST("/posts", ps.Register)
	api.GET("/posts", ps.GetAll)
	api.GET("/posts/:id", ps.GetById)
	api.PUT("/posts/:id", ps.UpdateById)
	api.PUT("/posts/:id/mark", ps.UpdateMarkById)
	api.DELETE("/posts/:id", ps.DeleteById)
}
