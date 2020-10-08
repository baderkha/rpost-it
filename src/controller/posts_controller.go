package controller

import (
	"comment-me/src/service"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	Service service.IPostService
	BaseController
}

func (p *PostController) CreatePost(c *gin.Context) {
	var postReq service.PostCreateRequestBody
	var postIdentity service.PostRequestQuery
	err := c.ShouldBindQuery(&postIdentity)
	if err != nil {
		p.GinInputError(c, err)
		return
	}
	err = c.ShouldBindJSON(&postReq)
	if err != nil {
		p.GinInputError(c, err)
		return
	}
	post, err := p.Service.CreatePostByAccount(&postIdentity, &postReq)
	if err != nil {
		p.GenerateResponseFromError(c, err)
		return
	}
	p.OK(c, post)
}

func (p *PostController) GetPostById(c *gin.Context) {
	post, err := p.Service.GetPostByid(c.Param("id"))
	if err != nil {
		p.GenerateResponseFromError(c, err)
		return
	}
	p.OK(c, post)
}
