package service

import (
	"comment-me/src/repository"
	"errors"
	"fmt"
)

type PostCreateRequest struct {
	Title       string `json:"title" binding:"required"`
	Text        string `json:"text" binding:"required"`
	CommunityId *uint  `json:"communityId" binding:"required" form:"communityId"`
}

type IPostService interface {
	CreatePostByAccount(accountId *uint, postReq *PostCreateRequest) (*repository.Post, error)
	GetPostByid(id string) (*repository.Post, error)
	GetPostsForCommunity(communityId string) *[]repository.Post
}

type PostService struct {
	Repo repository.IPostRepo
}

func (p *PostService) CreatePostByAccount(accountId *uint, postReq *PostCreateRequest) (*repository.Post, error) {
	var post repository.Post
	if accountId == nil {
		return nil, errors.New("400, Missing account Id to associate the post with")
	} else if postReq.CommunityId == nil {
		return nil, errors.New("400, Missing community Id to associate the post with")
	}

	post.AccountId = *accountId
	post.CommunityId = *postReq.CommunityId
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
