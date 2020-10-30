package controller

import (
	"rpost-it/src/service"

	"github.com/gin-gonic/gin"
)

// PostController : HTTP controller for the post
type PostController struct {
	Service service.IFacade
	BaseController
}

// CreatePost : Create a post
func (p *PostController) CreatePost(c *gin.Context) {
	uniqueID := c.Param("readableId")
	var postReq service.PostCreateRequestBody
	var postIdentity service.PostRequestQuery
	if uniqueID != "" {
		postIdentity.CommunityUniqueID = &uniqueID
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

// GetPostByID : Get a post by the internal id
func (p *PostController) GetPostByID(c *gin.Context) {
	post, err := p.Service.GetPostByID(c.Param("id"))
	if err != nil {
		p.GenerateResponseFromError(c, err)
		return
	}
	p.OK(c, post)
}

// GetPostsForCommunityByHumanReadibleID : Get the post by the internal id
func (p *PostController) GetPostsForCommunityByHumanReadibleID(c *gin.Context) {
	id := c.Param("readableId")
	if id == "" {
		p.BadRequest(c, "Id must be set")
		return
	}
	com, err := p.Service.GetPostsForCommunityByHumanReadibleID(id)
	if err != nil {
		p.GenerateResponseFromError(c, err)
		return
	}
	p.OK(c, com)
}
