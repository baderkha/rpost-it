package service

import (
	"errors"
	"fmt"
	"rpost-it/src/repository"
)

// PostCreateRequestBody :
type PostCreateRequestBody struct {
	Title string `json:"title" binding:"required"`
	Text  string `json:"text" binding:"required"`
}

// PostRequestQuery :
type PostRequestQuery struct {
	CommunityID       *uint   `json:"communityId" form:"communityId"`
	AccountID         *uint   `json:"accountId" binding:"required" form:"accountId"`
	CommunityUniqueID *string `json:"readableId" form:"readableId"`
}

// PostService : Posting service for posts :)
type PostService struct {
	Repo repository.IPostRepo
}

// CreatePostByAccount : Create a post by the account in question , given the post content
func (p *PostService) CreatePostByAccount(postIdentity *PostRequestQuery, postReq *PostCreateRequestBody) (*repository.Post, error) {
	var post repository.Post
	// guard check for post content
	if post.Text == "" || post.Title == "" {
		return nil, errors.New("400, Cannot have the Text or title empty")
	}

	// safley set stuff from input
	post.AccountId = *postIdentity.AccountID
	post.CommunityId = *postIdentity.CommunityID
	post.Text = postReq.Text
	post.Title = postReq.Title

	isCreated := p.Repo.CreatePost(&post)
	if !isCreated {
		return nil, errors.New("500, Could not create Post")
	}
	newPost, _ := p.Repo.FindByPostId(fmt.Sprintf("%d", post.ID))
	return newPost, nil
}

// GetPostByID : fetch a post by the unique internal id=
func (p *PostService) GetPostByID(id string) (*repository.Post, error) {
	if id == "" {
		return nil, errors.New("400, Expected id to be provided for the post")
	}
	post, isFound := p.Repo.FindByPostId(id)
	if !isFound {
		return nil, errors.New("404, No Post Found for this id")
	}
	return post, nil
}

// GetPostsForCommunity : Get all posts for a community
func (p *PostService) GetPostsForCommunity(communityID string) (*[]repository.Post, error) {
	posts := p.Repo.FindByCommunityId(communityID)
	return posts, nil
}

// GetPostsByCommunityID : get all posts for a community
func (p *PostService) GetPostsByCommunityID(communityID string) (*[]repository.Post, error) {
	posts := p.Repo.FindByCommunityId(communityID)
	return posts, nil
}
