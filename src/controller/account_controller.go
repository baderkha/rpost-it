package controller

import (
	"comment-me/src/service"
	"github.com/gin-gonic/gin"
)

type AccountController struct {
	BaseController
	Service service.IAccountService
}

func (a *AccountController) POSTAccount(c *gin.Context) {
	var req service.RegistrationDetails
	err := c.ShouldBindJSON(&req)
	if err != nil {
		a.GinInputError(c, err)
		return
	}
	acc, err := a.Service.RegisterAccountAndUser(&req)
	if err != nil {
		a.GenerateResponseFromError(c, err)
		return
	}
	a.Created(c, &acc)
}

func (a *AccountController) POSTAccountJWT(c *gin.Context) {
	var req service.LoginDetails
	err := c.ShouldBindJSON(&req)
	if err != nil {

		a.GinInputError(c, err)
		return
	}
	jwt, err := a.Service.LoginAccount(&req)
	if err != nil {
		a.GenerateResponseFromError(c, err)
		return
	}
	a.Created(c, &jwt)
}
