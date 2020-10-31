package controller

import (
	"rpost-it/src/service"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	stdMsg = "Your JWT is invalid"
)

// MiddleWareController : this will house all middle ware functionality for http
type MiddleWareController struct {
	Service service.IFacade
	BaseController
}

// VerifyJWTToken : Verify that a token is legit yo
func (mc *MiddleWareController) VerifyJWTToken(c *gin.Context) {
	// fetch the token from header
	auth := c.GetHeader("Authorization")
	if !strings.Contains(auth, "Bearer ") {
		mc.UnAuthorized(c, stdMsg)
		return
	}

	// get the non bearer part
	token := strings.Split(auth, "Bearer ")[0]

	err := mc.Service.ValidateJWT(token)
	if err != nil {
		mc.GenerateResponseFromError(c, err)
		return
	}

	c.Next()
}
