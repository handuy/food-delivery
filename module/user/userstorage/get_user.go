package userstorage

import "food-delivery/module/user/usermodel"

func(s *store) GetUser(email string) (usermodel.User, error) {
	var user usermodel.User
	err := s.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}