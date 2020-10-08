package repository

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Text        string // body content of the post
	Title       string `gorm:"type:varchar(255)"` // title of the post
	AccountId   uint   // account id that posted this
	Likes       uint   // total likes
	CommunityId uint   `gorm:"community_id"` // the community id
}

type IPostRepo interface {
	FindByPostId(id string) (*Post, bool)
	FindByAccountId(accountId string) *[]Post
	FindAlL() *[]Post
	FindByCommunityId(communityId string) *[]Post
	CreatePost(post *Post) bool
	DeletePostByPostId(post *Post) bool
	UpdatePost(Post *Post) bool
}

type PostRepo struct {
	BaseRepo
}

func (p *PostRepo) FindByPostId(id string) (*Post, bool) {
	var post Post
	isFound := p.FindById(id, &post)
	return &post, isFound
}

func (p *PostRepo) FindAll() *[]Post {
	var posts []Post
	p.GetContext().Find(&posts)
	return &posts
}

func (p *PostRepo) CreatePost(post *Post) bool {
	return p.Create(post)
}

func (p *PostRepo) DeletePostByPostId(id string) bool {
	return p.DeleteById(id, &Post{})
}

func (p *PostRepo) UpdatePost(post *Post) bool {
	return p.Update(post)
}

func (p *PostRepo) FindByAccountId(accountId string) *[]Post {
	var posts []Post
	p.
		GetContext().
		Where("account_id=?", accountId).
		Find(&posts)
	return &posts
}
