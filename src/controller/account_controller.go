package controller

import (
	"rpost-it/src/service"

	"github.com/gin-gonic/gin"
)

//AccountController : HTTP Account Controller
type AccountController struct {
	BaseController
	Service service.IFacade
}

// POSTAccount : Create Account
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

// POSTAccountJWT : Create a JWT for registered account
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

// POSTAccountJWTRefresh : refresh a jwt for an account Id
func (a *AccountController) POSTAccountJWTRefresh(c *gin.Context) {
	accountId := c.Query("accountId")
	if accountId == "" {
		a.BadRequest(c, "Expecting account id in the query paramters")
		return
	}
	jwt, err := a.Service.RefreshJWTToken(accountId)
	if err != nil {
		a.GenerateResponseFromError(c, err)
		return
	}
	a.Created(c, &jwt)
}

// GetAccountInfoByJWT : Fetch the account information from valid jwt token
func (a *AccountController) GetAccountInfoByJWT(c *gin.Context) {
	jwt := c.GetHeader("Authorization")
	if jwt == "" {
		a.BadRequest(c, "Expected a bearer token for this route")
		return
	}
	acc, err := a.Service.GetAccountInfoByJWT(jwt)
	if err != nil {
		a.GenerateResponseFromError(c, err)
		return
	}
	a.OK(c, acc)
}
