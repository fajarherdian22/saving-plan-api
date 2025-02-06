package controller

import (
	"github.com/fajarherdian22/saving-plan-api/exception"
	"github.com/fajarherdian22/saving-plan-api/helper"
	"github.com/fajarherdian22/saving-plan-api/service"
	"github.com/fajarherdian22/saving-plan-api/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type CustomerController struct {
	UserService *service.UserServiceImpl
	Validate    *validator.Validate
}

func NewUserController(UserService *service.UserServiceImpl, validate *validator.Validate) *CustomerController {
	return &CustomerController{
		UserService: UserService,
		Validate:    validate,
	}
}

func (controller *CustomerController) CreateUser(c *gin.Context) {
	var req web.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		exception.ErrorHandler(c, err)
		return
	}

	if err := controller.Validate.Struct(req); err != nil {
		exception.ErrorHandler(c, err)
		return
	}

	arg := web.CreateUserPayload(req)

	payload, err := controller.UserService.CreateUser(c.Request.Context(), arg)
	if err != nil {
		exception.ErrorHandler(c, err)
		return
	}

	WebResponse := web.WebResponse{
		Code:   200,
		Data:   payload,
		Status: "OK",
	}

	helper.HandleEncodeWriteJson(c, WebResponse)
}

func (controller *CustomerController) GetUser(c *gin.Context) {
	type GetUserReq struct {
		Email string `json:"email" binding:"required,email"`
	}
	var req GetUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		exception.ErrorHandler(c, err)
		return
	}

	if err := controller.Validate.Struct(req); err != nil {
		exception.ErrorHandler(c, err)
		return
	}

	payload, err := controller.UserService.GetUser(c.Request.Context(), req.Email)
	if err != nil {
		exception.ErrorHandler(c, err)
		return
	}

	WebResponse := web.WebResponse{
		Code:   200,
		Data:   payload,
		Status: "OK",
	}

	helper.HandleEncodeWriteJson(c, WebResponse)
}
