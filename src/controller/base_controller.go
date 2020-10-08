package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type BaseResponse struct {
	IsError  bool
	Message  string
	Resource interface{}
}

type BaseController struct {
}

func (b *BaseController) OK(c *gin.Context, resource interface{}) {
	c.JSON(http.StatusOK, BaseResponse{
		IsError:  false,
		Message:  "OK",
		Resource: resource,
	})
}

func (b *BaseController) GinInputError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusBadRequest, BaseResponse{
		IsError:  true,
		Message:  "BAD INPUT =>" + err.Error(),
		Resource: nil,
	})
}

func (b *BaseController) Created(c *gin.Context, resource interface{}) {
	c.JSON(http.StatusCreated, BaseResponse{
		IsError:  false,
		Message:  "Created Resource",
		Resource: resource,
	})
}

func (b *BaseController) NotFound(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusNotFound, BaseResponse{
		IsError:  true,
		Message:  message,
		Resource: nil,
	})
}

func (b *BaseController) InternalServerError(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, BaseResponse{
		IsError:  true,
		Message:  message,
		Resource: nil,
	})
}

func (b *BaseController) BadRequest(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, BaseResponse{
		IsError:  true,
		Message:  message,
		Resource: nil,
	})
}

func (b *BaseController) UnAuthorized(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, BaseResponse{
		IsError:  true,
		Message:  "UNAUTHORIZED => " + message,
		Resource: nil,
	})
}

func (b *BaseController) GenerateResponseFromError(c *gin.Context, err error) {
	message := err.Error()
	messageAr := strings.Split(message, ",")
	message = messageAr[0]
	messageCode, err := strconv.Atoi(message)
	if err != nil {
		b.InternalServerError(c, err.Error())
		return
	}
	switch messageCode {
	case http.StatusUnauthorized:
		b.UnAuthorized(c, messageAr[1])
		return
	case http.StatusNotFound:
		b.NotFound(c, messageAr[1])
		return
	case http.StatusBadRequest:
		b.BadRequest(c, messageAr[1])
		return
	default:
		b.InternalServerError(c, message)
		return
	}
}
