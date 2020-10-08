package util

import (
	"golang.org/x/crypto/bcrypt"
	//"time"
)

// Password Interface
type IPassword interface {
	// Generate a Hash for the password that's safley Storable
	HashPassword(password string, hashStrength int) (string, error)
	// Check if the password input matches the hash expected
	CheckPasswordHash(password string, hash string) bool
}

// Implements IPassword interface using bcrypt as an algorithm
type PasswordBcrypt struct {
}

func (p *PasswordBcrypt) HashPassword(password string, hashStrength int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), hashStrength)
	return string(bytes), err
}

func (p *PasswordBcrypt) CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
