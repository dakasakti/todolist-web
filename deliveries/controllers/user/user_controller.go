package controllers

import (
	"github.com/dakasakti/postingan/deliveries/helpers"
	"github.com/dakasakti/postingan/entities"
	us "github.com/dakasakti/postingan/services/user"
	"github.com/dakasakti/postingan/services/validation"

	"github.com/labstack/echo/v4"
)

type userController struct {
	Us us.UserService
	Vs validation.Validation
}

func NewUserController(us us.UserService, vs validation.Validation) *userController {
	return &userController{Us: us, Vs: vs}
}

func (uc *userController) Register(ctx echo.Context) error {
	var data entities.UserRequest

	err := ctx.Bind(&data)
	if err != nil {
		return ctx.JSON(400, helpers.ResponseJSON{
			Status:  400,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = uc.Vs.Validate(data)
	if err != nil {
		return ctx.JSON(400, helpers.ResponseJSON{
			Status:  400,
			Message: "input data is not valid",
			Data:    validation.MessageValidate(err),
		})
	}

	err = uc.Us.Register(data)
	if err != nil {
		return ctx.JSON(400, helpers.ResponseJSON{
			Status:  400,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(201, helpers.ResponseJSON{
		Status:  201,
		Message: "successfully created",
		Data:    nil,
	})
}

func (uc *userController) Login(ctx echo.Context) error {
	var data entities.LoginRequest

	err := ctx.Bind(&data)
	if err != nil {
		return ctx.JSON(400, helpers.ResponseJSON{
			Status:  400,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = uc.Vs.Validate(data)
	if err != nil {
		return ctx.JSON(400, helpers.ResponseJSON{
			Status:  400,
			Message: "input data is not valid",
			Data:    validation.MessageValidate(err),
		})
	}

	token, err := uc.Us.Login(data)
	if err != nil {
		return ctx.JSON(401, helpers.ResponseJSON{
			Status:  401,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(200, helpers.ResponseJSON{
		Status:  200,
		Message: "successfully logged in",
		Data:    token,
	})
}
