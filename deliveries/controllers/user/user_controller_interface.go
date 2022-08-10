package controllers

import "github.com/labstack/echo/v4"

type UserController interface {
	Register(ctx echo.Context) error
	Login(ctx echo.Context) error
	Profile(ctx echo.Context) error
}
