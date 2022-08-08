package client

import "github.com/labstack/echo/v4"

type ClientController interface {
	GetAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Store(ctx echo.Context) error
	Edit(ctx echo.Context) error
	UpdateData(ctx echo.Context) error
	UpdateMark(ctx echo.Context) error

	// Auth
	Index(ctx echo.Context) error
	StoreAuth(ctx echo.Context) error
	LoginAuth(ctx echo.Context) error
	Logout(ctx echo.Context) error
}
