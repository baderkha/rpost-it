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
	CommunityId       *uint   `json:"communityId" form:"communityId"`
	AccountId         *uint   `json:"accountId" binding:"required" form:"accountId"`
	CommunityUniqueId *string `json:"readableId" form:"readableId"`
}

type IPostService interface {
	CreatePostByAccount(postIdentity *PostRequestQuery, postReq *PostCreateRequestBody) (*repository.Post, error)
	GetPostByid(id string) (*repository.Post, error)
	GetPostsForCommunity(communityId string) (*[]repository.Post, error)
	GetPostsForCommunityByUniqueId(uniqueId string) (*[]repository.Post, error)
}

type PostService struct {
	Repo             repository.IPostRepo
	AccountService   IAccountService
	CommunityService ICommunityService
}

func (p *PostService) CreatePostByAccount(postIdentity *PostRequestQuery, postReq *PostCreateRequestBody) (*repository.Post, error) {
	var post repository.Post

	// validation happens here
	if postIdentity.AccountId == nil {
		return nil, errors.New("400, Missing account Id to associate the post with")
	} else if postIdentity.CommunityId == nil && postIdentity.CommunityUniqueId == nil {
		return nil, errors.New("400, Missing community Id to associate the post with")
	} else if postIdentity.CommunityId == nil && postIdentity.CommunityUniqueId != nil {
		// fetch the community internal id
		com, err := p.CommunityService.FindCommunityByUniqueID(*postIdentity.CommunityUniqueId)
		if err != nil {
			return nil, err
		}
		postIdentity.CommunityId = &com.ID
	} else if postIdentity.CommunityId != nil && postIdentity.CommunityUniqueId != nil {
		return nil, errors.New("400, You cannot have both set set by unique id and the query parameter")
	}

	// verify it exists atleast
	_, err := p.CommunityService.FindCommunityByID(fmt.Sprint(*postIdentity.CommunityId))
	if err != nil {
		return nil, err
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

func (p *PostService) GetPostsForCommunity(communityId string) (*[]repository.Post, error) {
	posts := p.Repo.FindByCommunityId(communityId)
	return posts, nil
}

func (p *PostService) GetPostsForCommunityByUniqueId(unqiueId string) (*[]repository.Post, error) {
	community, err := p.CommunityService.FindCommunityByUniqueID(unqiueId)
	if err != nil {
		return nil, err
	}
	return p.GetPostsForCommunity(fmt.Sprint(community.ID))
}
