package service

import (
	"errors"
	"fmt"
	"rpost-it/src/repository"
)

// CreateCommunityBody  : Body we expect from the request for create
type CreateCommunityBody struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	UniqueID    string `json:"uniqueId"`
}

// CommunityIdentitiy : How to know who this commuinity belongs to
type CommunityIdentitiy struct {
	AccountOwner uint `form:"accountId" binding:"required"`
}

// CommunityIdentifier : The community by the unique human readable id
type CommunityIdentifier struct {
	UniqueID string `form:"uniqueId" binding:"required"`
}

// CommunityService : concrete implementation fo the community service
type CommunityService struct {
	Repo repository.ICommunityRepo
}

// CreateCommunity : Creates a COMMUNITY if one des not exist by the unique id
func (c *CommunityService) CreateCommunity(identity *CommunityIdentitiy, comBody *CreateCommunityBody) (*repository.Community, error) {

	// check the important fields
	if comBody.UniqueID == "" || comBody.Title == "" {
		return nil, errors.New("400, Unique id or title must be set")
	}

	// make sure we don't have one that exists , that will cause problems
	_, isAlreadyExists := c.Repo.FindCommunityByUniqueID(comBody.UniqueID)
	if isAlreadyExists {
		return nil, errors.New("400, Community Already Exists")
	}

	// define the model
	var community repository.Community
	community.AccountOwner = identity.AccountOwner
	community.Title = comBody.Title
	community.UniqueID = comBody.UniqueID
	community.Description = comBody.Description

	// incase it's not created there just freak out
	isCreated := c.Repo.CreateCommunity(&community)
	if !isCreated {
		return nil, errors.New("500, could not create the record")
	}

	// not expecting a bad return , so we can assume it will
	freshCom, _ := c.Repo.FindCommunityByID(fmt.Sprint(community.ID))
	return freshCom, nil
}

// FindCommunityByUniqueID : Get the community by the human readable id
func (c *CommunityService) FindCommunityByUniqueID(uniqueID string) (*repository.Community, error) {
	if uniqueID == "" {
		return nil, errors.New("400, Id should not be empty")
	}
	community, isFound := c.Repo.FindCommunityByUniqueID(uniqueID)
	if !isFound {
		return nil, errors.New("404, Could not Find a community for this identifier")
	}
	return community, nil
}

// FindCommunityByID : Finds the commuinity byu the primary key db id, this is the internal id
func (c *CommunityService) FindCommunityByID(id string) (*repository.Community, error) {
	if id == "" {
		return nil, errors.New("400, Id should not be empty")
	}
	community, isFound := c.Repo.FindCommunityByID(id)
	if !isFound {
		return nil, errors.New("404, Could not Find a community for this internal id")
	}
	return community, nil
}
