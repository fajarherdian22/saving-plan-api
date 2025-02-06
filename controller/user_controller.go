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

func NewCustomerController(UserService *service.UserServiceImpl, validate *validator.Validate) *CustomerController {
	return &CustomerController{
		UserService: UserService,
		Validate:    validate,
	}
}

func (controller *CustomerController) CreateCustomersUser(c *gin.Context) {
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
