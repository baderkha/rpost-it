package service

import (
	"comment-me/src/repository"
	"errors"
	"fmt"
)

type PostCreateRequestBody struct {
	Title string `json:"title" binding:"required"`
	Text  string `json:"text" binding:"required"`
}

type PostRequestQuery struct {
	CommunityId *uint `json:"communityId" binding:"required" form:"communityId"`
	AccountId   *uint `json:"accountId" binding:"required" form:"accountId"`
}

type IPostService interface {
	CreatePostByAccount(postIdentity *PostRequestQuery, postReq *PostCreateRequestBody) (*repository.Post, error)
	GetPostByid(id string) (*repository.Post, error)
	GetPostsForCommunity(communityId string) *[]repository.Post
}

type PostService struct {
	Repo           repository.IPostRepo
	AccountService IAccountService
}

func (p *PostService) CreatePostByAccount(postIdentity *PostRequestQuery, postReq *PostCreateRequestBody) (*repository.Post, error) {
	var post repository.Post
	if postIdentity.AccountId == nil {
		return nil, errors.New("400, Missing account Id to associate the post with")
	} else if postIdentity.CommunityId == nil {
		return nil, errors.New("400, Missing community Id to associate the post with")
	}

	isAccountExists := p.AccountService.ValidateAccountExists(fmt.Sprintf("%d", *postIdentity.AccountId))
	if !isAccountExists {
		return nil, errors.New("400, This Account Does not Exist")
	}
	post.AccountId = *postIdentity.AccountId
	post.CommunityId = *postIdentity.CommunityId
	post.Text = postReq.Text
	post.Title = postReq.Title

	isCreated := p.Repo.CreatePost(&post)
	if !isCreated {
		return nil, errors.New("500, Could not create Post")
	}
	newPost, _ := p.Repo.FindByPostId(fmt.Sprintf("%d", post.ID))
	return newPost, nil
}

func (p *PostService) GetPostByid(id string) (*repository.Post, error) {
	if id == "" {
		return nil, errors.New("400, Expected id to be provided for the post")
	}
	post, isFound := p.Repo.FindByPostId(id)
	if !isFound {
		return nil, errors.New("404, No Post Found for this id")
	}
	return post, nil
}

func (p *PostService) GetPostsForCommunity(communityId string) *[]repository.Post {
	posts := p.Repo.FindByCommunityId(communityId)
	return posts
}
