package repository

import (
	"github.com/chanpon2015/authentication/model"
	"github.com/jinzhu/gorm"
)

// UsersRepository is
type UsersRepository interface {
	FindByUserID(string) (*model.User, error)
	Insert(model.User) error
}

type usersRepositoryImpl struct {
	db *gorm.DB
}

// NewUsersRepository is
func NewUsersRepository(db *gorm.DB) UsersRepository {
	return &usersRepositoryImpl{db}
}

func (r *usersRepositoryImpl) FindByUserID(userID string) (*model.User, error) {
	user := model.User{}
	result := r.db.Find(&user, "user_id = ?", userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *usersRepositoryImpl) Insert(user model.User) error {
	return r.db.Save(&user).Error
}
