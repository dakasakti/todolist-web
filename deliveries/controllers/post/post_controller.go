package controllers

import (
	"github.com/dakasakti/todolist-web/deliveries/helpers"
	"github.com/dakasakti/todolist-web/deliveries/middlewares"
	"github.com/dakasakti/todolist-web/entities"
	ps "github.com/dakasakti/todolist-web/services/post"
	"github.com/dakasakti/todolist-web/services/validation"

	"github.com/labstack/echo/v4"
)

type postController struct {
	Ps ps.PostService
	Vs validation.Validation
}

func NewPostController(ps ps.PostService, vs validation.Validation) *postController {
	return &postController{Ps: ps, Vs: vs}
}

func (pc *postController) Register(ctx echo.Context) error {
	var data entities.PostRequest
	user_id := uint(middlewares.ExtractTokenUserId(ctx))

	err := ctx.Bind(&data)
	if err != nil {
		return ctx.JSON(400, helpers.ResponseJSON{
			Status:  400,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = pc.Vs.Validate(data)
	if err != nil {
		return ctx.JSON(400, helpers.ResponseJSON{
			Status:  400,
			Message: "input data is not valid",
			Data:    validation.MessageValidate(err),
		})
	}

	err = pc.Ps.Register(user_id, data)
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

func (pc *postController) GetAll(ctx echo.Context) error {
	data, err := pc.Ps.GetAll()
	if err != nil {
		return ctx.JSON(400, helpers.ResponseJSON{
			Status:  400,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if data == nil {
		return ctx.JSON(404, helpers.ResponseJSON{
			Status:  404,
			Message: "record not found",
			Data:    nil,
		})
	}

	return ctx.JSON(200, helpers.ResponseJSON{
		Status:  200,
		Message: "successfully retrieved",
		Data:    data,
	})
}

func (pc *postController) GetById(ctx echo.Context) error {
	param := ctx.Param("id")

	id, err := pc.Ps.CheckParamId(param)
	if err != nil {
		return ctx.JSON(400, helpers.ResponseJSON{
			Status:  400,
			Message: "param only use id (number)",
			Data:    nil,
		})
	}

	data, err := pc.Ps.GetById(id)
	if err != nil {
		return ctx.JSON(404, helpers.ResponseJSON{
			Status:  404,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(200, helpers.ResponseJSON{
		Status:  200,
		Message: "successfully retrieved",
		Data:    data,
	})
}

func (pc *postController) UpdateById(ctx echo.Context) error {
	var data entities.PostUpdateRequest
	param := ctx.Param("id")
	user_id := uint(middlewares.ExtractTokenUserId(ctx))

	id, err := pc.Ps.CheckParamId(param)
	if err != nil {
		return ctx.JSON(400, helpers.ResponseJSON{
			Status:  400,
			Message: "param only use id (number)",
			Data:    nil,
		})
	}

	result, err := pc.Ps.CheckUser(id, user_id)
	if err != nil {
		return ctx.JSON(403, helpers.ResponseJSON{
			Status:  403,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(400, helpers.ResponseJSON{
			Status:  400,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = pc.Vs.Validate(data)
	if err != nil {
		return ctx.JSON(400, helpers.ResponseJSON{
			Status:  400,
			Message: "input data is not valid",
			Data:    validation.MessageValidate(err),
		})
	}

	err = pc.Ps.UpdateById(result.ID, data)
	if err != nil {
		return ctx.JSON(400, helpers.ResponseJSON{
			Status:  400,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(200, helpers.ResponseJSON{
		Status:  200,
		Message: "successfully updated",
		Data:    nil,
	})
}

func (pc *postController) UpdateMarkById(ctx echo.Context) error {
	param := ctx.Param("id")
	user_id := uint(middlewares.ExtractTokenUserId(ctx))

	id, err := pc.Ps.CheckParamId(param)
	if err != nil {
		return ctx.JSON(400, helpers.ResponseJSON{
			Status:  400,
			Message: "param only use id (number)",
			Data:    nil,
		})
	}

	result, err := pc.Ps.CheckUser(id, user_id)
	if err != nil {
		return ctx.JSON(403, helpers.ResponseJSON{
			Status:  403,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = pc.Ps.UpdateMarkById(result.ID)
	if err != nil {
		return ctx.JSON(400, helpers.ResponseJSON{
			Status:  400,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(200, helpers.ResponseJSON{
		Status:  200,
		Message: "successfully updated",
		Data:    nil,
	})
}

func (pc *postController) DeleteById(ctx echo.Context) error {
	param := ctx.Param("id")
	user_id := uint(middlewares.ExtractTokenUserId(ctx))

	id, err := pc.Ps.CheckParamId(param)
	if err != nil {
		return ctx.JSON(400, helpers.ResponseJSON{
			Status:  400,
			Message: "param only use id (number)",
			Data:    nil,
		})
	}

	result, err := pc.Ps.CheckUser(id, user_id)
	if err != nil {
		return ctx.JSON(403, helpers.ResponseJSON{
			Status:  403,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = pc.Ps.DeleteById(result.ID)
	if err != nil {
		return ctx.JSON(400, helpers.ResponseJSON{
			Status:  400,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(200, helpers.ResponseJSON{
		Status:  200,
		Message: "successfully deleted",
		Data:    nil,
	})
}
