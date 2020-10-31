package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// BaseResponse : this is the base response the http methods return
type BaseResponse struct {
	IsError  bool
	Message  string
	Resource interface{}
}

// BaseController our baseController
type BaseController struct {
}

// OK : 200 response with data
func (b *BaseController) OK(c *gin.Context, resource interface{}) {
	c.JSON(http.StatusOK, BaseResponse{
		IsError:  false,
		Message:  "OK",
		Resource: resource,
	})
}

// GinInputError : use this for Gin errors
func (b *BaseController) GinInputError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusBadRequest, BaseResponse{
		IsError:  true,
		Message:  "BAD INPUT =>" + err.Error(),
		Resource: nil,
	})
}

// Created : 201 with created resource
func (b *BaseController) Created(c *gin.Context, resource interface{}) {
	c.JSON(http.StatusCreated, BaseResponse{
		IsError:  false,
		Message:  "Created Resource",
		Resource: resource,
	})
}

// NotFound : 404 , with a custom message
func (b *BaseController) NotFound(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusNotFound, BaseResponse{
		IsError:  true,
		Message:  message,
		Resource: nil,
	})
}

// InternalServerError : 500 , with an internal message
func (b *BaseController) InternalServerError(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, BaseResponse{
		IsError:  true,
		Message:  message,
		Resource: nil,
	})
}

// BadRequest : 400 with a custom message
func (b *BaseController) BadRequest(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, BaseResponse{
		IsError:  true,
		Message:  message,
		Resource: nil,
	})
}

// UnAuthorized : 401 with a custom message
func (b *BaseController) UnAuthorized(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, BaseResponse{
		IsError:  true,
		Message:  "UNAUTHORIZED => " + message,
		Resource: nil,
	})
}

// GenerateResponseFromError : take an error object from some service that follows the following pattern
// "<<(UINT)CODE>>, <<(STRING)some_message>>"
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
