package service

import (
	"errors"
	"fmt"
	"rpost-it/src/repository"
	"rpost-it/src/util"

	"gorm.io/gorm"
)

// RegistrationResponse : things to send back on registration
type RegistrationResponse struct {
	JWT     *JWT
	Account *repository.Account
}

// a little check to make sure we impl the ifc
var facade IFacade = &Facade{}

// singleton pointer
var facadePtr *Facade

// IFacade : Abstract Facade for the service
type IFacade interface {
	// Registers an account and a user together , and generates a JWT
	RegisterAccountAndUser(r *RegistrationDetails) (*RegistrationResponse, error)
	// Validate Login to account and generate a jwt
	LoginAccount(l *LoginDetails) (*JWT, error)
	// fetch account information for a valid jwt token
	GetAccountInfoByJWT(JWTBearer string) (*repository.Account, error)
	// Creates a community by a user
	CreateCommunity(identity *CommunityIdentitiy, comBody *CreateCommunityBody) (*repository.Community, error)
	// Fetch a community by a human readible id
	FindCommunityByUniqueID(uniqueID string) (*repository.Community, error)
	// fetches a community by the internal id
	FindCommunityByID(ID string) (*repository.Community, error)
	// GetPostByID : fetch a post by the unique internal id
	GetPostByID(id string) (*repository.Post, error)
	// GetPostsForCommunity : Get all posts for a community
	GetPostsForCommunity(communityID string) (*[]repository.Post, error)
	// Create a post by a valid account
	CreatePostByAccount(postIdentity *PostRequestQuery, postReq *PostCreateRequestBody) (*repository.Post, error)
	// get posts for a human readible community id
	GetPostsForCommunityByHumanReadibleID(unqiueID string) (*[]repository.Post, error)
}

// Facade : Facade service implementation , this will allow us to orchestrate difficult logic that needs to talk to each other
type Facade struct {
	AccountSvc AccountService
	CommSvc    CommunityService
	PostSvc    PostService
	UserSvc    UserService
}

// MakeFacade : make the service facade
func MakeFacade(db *gorm.DB,
	jwtHelper util.IJwtHelper,
	jwtValidityMinutes int64,
	passwordHashStrengh uint,
	passwordHelper util.IPassword,
) *Facade {
	if facadePtr != nil {
		return facadePtr
	}
	facadePtr = &Facade{
		AccountSvc: AccountService{
			JWTHelper:              jwtHelper,
			JWTValidityMinutes:     jwtValidityMinutes,
			PassWordHashedStrength: passwordHashStrengh,
			PasswordHelper:         passwordHelper,
			Repo: &repository.AccountRepo{
				BaseRepo: *repository.MakeBaseRepo(db),
			},
		},
		CommSvc: CommunityService{
			Repo: &repository.CommunityRepo{
				BaseRepo: *repository.MakeBaseRepo(db),
			},
		},
		PostSvc: PostService{
			Repo: &repository.PostRepo{
				BaseRepo: *repository.MakeBaseRepo(db),
			},
		},
		UserSvc: UserService{
			Repo: &repository.UserRepo{
				BaseRepo: *repository.MakeBaseRepo(db),
			},
		},
	}
	return facadePtr
}

// RegisterAccountAndUser : Registers an account and a user together
func (f *Facade) RegisterAccountAndUser(r *RegistrationDetails) (*RegistrationResponse, error) {
	acc, err := f.AccountSvc.RegisterAccount(r)
	if err != nil {
		return nil, err
	}
	user, err := f.UserSvc.RegisterUser(&UserRegistrationDetails{
		AccountID:   fmt.Sprintf("%d", acc.ID),
		DateOfBirth: r.DateOfBirth,
		FirstName:   r.FirstName,
		LastName:    r.LastName,
	})
	if err != nil {
		// delete the account on profile failure
		_ = f.AccountSvc.DeleteServiceByIDInternalOnly(fmt.Sprintf("%d", acc.ID))
		return nil, err
	}

	acc.User = *user
	jwt, _ := f.AccountSvc.generateJWTForValidAccount(acc) // not expecting errors :)

	return &RegistrationResponse{
		Account: acc,
		JWT:     jwt,
	}, nil
}

// LoginAccount : Login to an account and generate a jwt
func (f *Facade) LoginAccount(l *LoginDetails) (*JWT, error) {
	return f.AccountSvc.LoginAccount(l)
}

// GetAccountInfoByJWT : Fetch account information by jwt token
func (f *Facade) GetAccountInfoByJWT(JWTBearer string) (*repository.Account, error) {
	return f.AccountSvc.GetAccountInfoByJWT(JWTBearer)
}

// CreateCommunity : proxies the create community method , and does an account check
func (f *Facade) CreateCommunity(identity *CommunityIdentitiy, comBody *CreateCommunityBody) (*repository.Community, error) {
	// ensure we have account
	accountExists := f.AccountSvc.ValidateAccountExists(fmt.Sprint(identity.AccountOwner))
	if !accountExists {
		return nil, errors.New("400, This account does not exist")
	}
	return f.CommSvc.CreateCommunity(identity, comBody)
}

// CreatePostByAccount :  this will verify account , and community before allowing proxy of creating the post
func (f *Facade) CreatePostByAccount(postIdentity *PostRequestQuery, postReq *PostCreateRequestBody) (*repository.Post, error) {

	// prechecks , kind of looks ugly , needs clean up
	if postIdentity.AccountID == nil {
		return nil, errors.New("400, Missing account Id to associate the post with")
	} else if postIdentity.CommunityID == nil && postIdentity.CommunityUniqueID == nil {
		return nil, errors.New("400, Missing community Id to associate the post with")
	} else if postIdentity.CommunityID == nil && postIdentity.CommunityUniqueID != nil {
		// fetch the community internal id
		com, err := f.CommSvc.FindCommunityByUniqueID(*postIdentity.CommunityUniqueID)
		if err != nil {
			return nil, err
		}
		postIdentity.CommunityID = &com.ID
	} else if postIdentity.CommunityID != nil && postIdentity.CommunityUniqueID != nil {
		return nil, errors.New("400, You cannot have both set set by unique id  and the internal id pick one :0")
	}

	// verify community it belongs to exists atleast
	_, err := f.CommSvc.FindCommunityByID(fmt.Sprint(*postIdentity.CommunityID))
	if err != nil {
		return nil, err
	}

	// verify we have an account
	isAccountExists := f.AccountSvc.ValidateAccountExists(fmt.Sprintf("%d", *postIdentity.AccountID))
	if !isAccountExists {
		return nil, errors.New("400, This Account Does not Exist")
	}

	// finally
	return f.PostSvc.CreatePostByAccount(postIdentity, postReq)
}

// FindCommunityByUniqueID : Finds a community by a community id that's human readible
func (f *Facade) FindCommunityByUniqueID(uniqueID string) (*repository.Community, error) {
	return f.CommSvc.FindCommunityByUniqueID(uniqueID)
}

// FindCommunityByID : Finds the commuinity byu the primary key db id, this is the internal id
func (f *Facade) FindCommunityByID(id string) (*repository.Community, error) {
	return f.CommSvc.FindCommunityByID(id)
}

// GetPostByID : Find a post by post internal id
func (f *Facade) GetPostByID(id string) (*repository.Post, error) {
	return f.PostSvc.GetPostByID(id)
}

// GetPostsForCommunityByHumanReadibleID :fetches the posts by the human readible id
func (f *Facade) GetPostsForCommunityByHumanReadibleID(unqiueID string) (*[]repository.Post, error) {
	community, err := f.CommSvc.FindCommunityByUniqueID(unqiueID)
	if err != nil {
		return nil, err
	}
	posts, err := f.PostSvc.GetPostsByCommunityID(fmt.Sprintf("%d", community.ID))
	return posts, err
}

// GetPostsForCommunity : Return posts for a community id , this is the inernal id for community
func (f *Facade) GetPostsForCommunity(communityID string) (*[]repository.Post, error) {
	// check if we have a community
	_, err := f.CommSvc.FindCommunityByID(communityID)
	if err != nil {
		return nil, err
	}
	return f.PostSvc.GetPostsByCommunityID(communityID)
}
