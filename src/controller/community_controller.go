package controller

import (
	"rpost-it/src/service"

	"github.com/gin-gonic/gin"
)

// CommunityController : Client facing community controller , takes gin operations
type CommunityController struct {
	BaseController
	Service service.IFacade
}

// CreateCommunity : Creates a new community !
func (cc *CommunityController) CreateCommunity(c *gin.Context) {
	var body service.CreateCommunityBody
	var identifier service.CommunityIdentitiy
	// bind the stuff needed
	err := c.ShouldBindQuery(&identifier)
	if err != nil {
		cc.GinInputError(c, err)
		return
	}
	err = c.ShouldBindJSON(&body)
	if err != nil {
		cc.GinInputError(c, err)
		return
	}

	newCom, err := cc.Service.CreateCommunity(&identifier, &body)
	if err != nil {
		cc.GenerateResponseFromError(c, err)
		return
	}
	cc.OK(c, newCom)
}

// GetByHumanReadibleID : Finds the community by the uniqueid , ie codwarzone
func (cc *CommunityController) GetByHumanReadibleID(c *gin.Context) {
	id := c.Param("readableId")
	if id == "" {
		cc.BadRequest(c, "Id must be set")
		return
	}
	com, err := cc.Service.FindCommunityByUniqueID(id)
	if err != nil {
		cc.GenerateResponseFromError(c, err)
		return
	}
	cc.OK(c, com)
}

// GetByInternalID : Finds the community by the internal db id , this is number usually or uuid
func (cc *CommunityController) GetByInternalID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		cc.BadRequest(c, "Id must be set")
		return
	}
	com, err := cc.Service.FindCommunityByID(id)
	if err != nil {
		cc.GenerateResponseFromError(c, err)
		return
	}
	cc.OK(c, com)
}
