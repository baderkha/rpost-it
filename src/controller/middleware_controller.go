package controller

import (
	"fmt"
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
func (mc MiddleWareController) VerifyJWTToken(c *gin.Context) {

	// fetch the token from header
	auth := c.GetHeader("Authorization")
	if !strings.Contains(auth, "Bearer ") {
		mc.UnAuthorized(c, "You're missing a token for this route => must be Bearer <<someToken>>")
		return
	} else if c.Query("accountId") == "" {
		mc.UnAuthorized(c, "You're missing an account id associated with this request. Query parameter accountId")
		return
	}

	// get the non bearer part
	token := strings.Split(auth, "Bearer ")[1]
	acc, err := mc.Service.ValidateJWT(token)
	if err != nil {
		mc.GenerateResponseFromError(c, err)
		return
	}

	// make sure jwt acc id matches the quer header
	if fmt.Sprintf("%d", acc.ID) != c.Query("accountId") {
		mc.UnAuthorized(c, stdMsg)
		return
	}
	c.Next()
}
