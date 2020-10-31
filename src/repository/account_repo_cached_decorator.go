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

// AccountRepositoryCachedDecorator : repository that follows a decorator pattern
type AccountRepositoryCachedDecorator struct {
	SQLAccountRepo IAccountRepo     // this must be an sql account repo
	CachingRepo    ICacheRepository // any type of caching layer , in memory , file , memory , s3 heck you name it :)
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
	acc, isFound := a.SQLAccountRepo.FindByAvatarId(avatarId)
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
	acc, isFound := a.SQLAccountRepo.FindByAvatarIdOrByEmail(avatarId, email)
	cacheVal := cacheAccountItem{
		Item:    acc,
		IsFound: isFound,
	}
	a.CachingRepo.Set(key, cacheVal, defaultAccountCacheTimeSeconds) // add to cache for next time
	return acc, isFound
}

// CreateAccount : Create account via sql impl just a proxy method
func (a *AccountRepositoryCachedDecorator) CreateAccount(acc *Account) bool {
	return a.SQLAccountRepo.CreateAccount(acc)
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
	acc, isFound := a.SQLAccountRepo.FindByAvatarId(id)
	cacheVal := cacheAccountItem{
		Item:    acc,
		IsFound: isFound,
	}
	a.CachingRepo.Set(key, cacheVal, defaultAccountCacheTimeSeconds) // add to cache for next time
	return acc, isFound
}

// DeleteAccountById : proxies the normal delete from sql
func (a *AccountRepositoryCachedDecorator) DeleteAccountById(id string) bool {
	return a.SQLAccountRepo.DeleteAccountById(id)
}
