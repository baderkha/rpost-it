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
	uniqueId := c.Param("readableId")
	var postReq service.PostCreateRequestBody
	var postIdentity service.PostRequestQuery
	if uniqueId != "" {
		postIdentity.CommunityUniqueId = &uniqueId
	}
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

func (p *PostController) GetPostsByUniqueCommunityId(c *gin.Context) {
	id := c.Param("readableId")
	if id == "" {
		p.BadRequest(c, "Id must be set")
		return
	}
	com, err := p.Service.GetPostsForCommunityByUniqueId(id)
	if err != nil {
		p.GenerateResponseFromError(c, err)
		return
	}
	p.OK(c, com)
}
