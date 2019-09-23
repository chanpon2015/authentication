package auth

import (
	"github.com/chanpon2015/authentication/model"
	"github.com/chanpon2015/authentication/repository"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Add is
func Add(db *gorm.DB, userID, password, name string) error {
	usersRepo := repository.NewUsersRepository(db)
	hash, err := hash(password)
	if err != nil {
		return err
	}
	user := model.User{
		UserID:   userID,
		Password: hash,
		Name:     name,
	}
	return usersRepo.Insert(user)
}

func hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
