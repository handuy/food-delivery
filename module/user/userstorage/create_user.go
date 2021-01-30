package userstorage

import (
	"errors"
	"food-delivery/module/user/usermodel"

	"gorm.io/gorm"
)

func (s *store) CheckEmailExist(email string) error {
	var user usermodel.User
	err := s.db.Where("email = ?", email).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return err
	}

	if err != nil {
		return err
	}

	if user.Id == "" {
		return errors.New("User ID không hợp lệ")
	}

	return nil
}

func (s *store) Create(user usermodel.User) (string, error) {
	err := s.db.Create(&user).Error
	if err != nil {
		return "", err
	}

	return user.Id, nil
}
