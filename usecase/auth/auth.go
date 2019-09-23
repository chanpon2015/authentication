package auth

import (
	"github.com/chanpon2015/authentication/repository"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Authentication is
func Authentication(db *gorm.DB, userID, password string) bool {
	usersRepo := repository.NewUsersRepository(db)
	user, err := usersRepo.FindByUserID(userID)
	if err != nil {
		return false
	}
	return compare(user.Password, password)
}

func compare(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
