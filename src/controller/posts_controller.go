package controller

import (
	"comment-me/src/service"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	Service service.IPostService
}

func (p *PostController) CreatePost(c *gin.Context) {
	
	p.Service.CreatePostByAccount()
}
