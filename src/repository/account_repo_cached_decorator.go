package repository

import (
	"fmt"
)

// implements the cached repository
const (
	basePrefixAccountCache         string = "account"
	defaultAccountCacheTimeSeconds uint   = 15
)

// cacheAccountItem : a cached item , singular item
type cacheAccountItem struct {
	Item    *Account
	IsFound bool
}

// Why do we need a decorator ?

// we need something that can masquerade as the audienceRepo
// and implement the methods fromt he parent interface
// with extra logic to use our caching layer

// this will safe our service from writing extra repetitve code that says
// if in cache , pull else , don't and so on ....if else if else

// so for the service it can just call that method once and from its perspective
// it's what it's expecting since it's using the repository interface

// also the icing to the cake , this is easily just unit testable no need for intetgration test
//  by just mocking the two dependencies and making sure
//  the key and values are what are passed !

// AccountRepositoryCachedDecorator : repository that follows a decorator pattern
type AccountRepositoryCachedDecorator struct {
	PersistentAccountRepo IAccountRepo     // this must be an sql account repo , or mongo or something
	CachingRepo           ICacheRepository // any type of caching layer , in memory , file , memory , s3 heck you name it :)
}

// FindByAvatarId : find by cache if hit return account by avatar id , else use sql
func (a *AccountRepositoryCachedDecorator) FindByAvatarId(avatarId string) (*Account, bool) {
	var key string = fmt.Sprintf(
		"%s::av_id=%s",
		basePrefixAccountCache,
		avatarId,
	)

	// make a check if we have something in cache
	cacheItem, isFound := a.CachingRepo.Get(key)
	if isFound {
		rtrn, ok := cacheItem.(*cacheAccountItem)
		if ok {
			return rtrn.Item, rtrn.IsFound
		}
	}

	// if we find the account repo
	acc, isFound := a.PersistentAccountRepo.FindByAvatarId(avatarId)
	cacheVal := cacheAccountItem{
		Item:    acc,
		IsFound: isFound,
	}
	a.CachingRepo.Set(key, cacheVal, defaultAccountCacheTimeSeconds) // add to cache for next time
	return acc, isFound
}

// FindByAvatarIdOrByEmail : find by cache if hit return account by avatar id or email, else use sql
func (a *AccountRepositoryCachedDecorator) FindByAvatarIdOrByEmail(avatarId string, email string) (*Account, bool) {
	var key string = fmt.Sprintf(
		"%s::av_id=%s||email=%s",
		basePrefixAccountCache,
		avatarId,
		email,
	)

	// make a check if we have something in cache
	cacheItem, isFound := a.CachingRepo.Get(key)
	if isFound {
		rtrn, ok := cacheItem.(*cacheAccountItem)
		if ok {
			return rtrn.Item, rtrn.IsFound
		}
	}

	// if we find the account repo
	acc, isFound := a.PersistentAccountRepo.FindByAvatarIdOrByEmail(avatarId, email)
	cacheVal := cacheAccountItem{
		Item:    acc,
		IsFound: isFound,
	}
	a.CachingRepo.Set(key, cacheVal, defaultAccountCacheTimeSeconds) // add to cache for next time
	return acc, isFound
}

// CreateAccount : Create account via sql impl just a proxy method
func (a *AccountRepositoryCachedDecorator) CreateAccount(acc *Account) bool {
	return a.PersistentAccountRepo.CreateAccount(acc)
}

// FindByAccountId : find an account via the cache layer , if found then cool , otherwise fetch from sql
func (a *AccountRepositoryCachedDecorator) FindByAccountId(id string) (*Account, bool) {
	var key string = fmt.Sprintf(
		"%s::id=%s",
		basePrefixAccountCache,
		id,
	)
	// make a check if we have something in cache
	cacheItem, isFound := a.CachingRepo.Get(key)
	if isFound {
		rtrn, ok := cacheItem.(*cacheAccountItem)
		if ok {
			return rtrn.Item, rtrn.IsFound
		}
	}

	// if we find the account repo
	acc, isFound := a.PersistentAccountRepo.FindByAvatarId(id)
	cacheVal := cacheAccountItem{
		Item:    acc,
		IsFound: isFound,
	}
	a.CachingRepo.Set(key, cacheVal, defaultAccountCacheTimeSeconds) // add to cache for next time
	return acc, isFound
}

// DeleteAccountById : proxies the normal delete from sql
func (a *AccountRepositoryCachedDecorator) DeleteAccountById(id string) bool {
	return a.PersistentAccountRepo.DeleteAccountById(id)
}
