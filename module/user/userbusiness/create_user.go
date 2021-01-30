package userbusiness

import (
	"food-delivery/common"
	"food-delivery/module/user/usermodel"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type CreateUserStore interface {
	CheckEmailExist(email string) error
	Create(user usermodel.User) (string, error)
}

type createUserBusiness struct {
	store CreateUserStore
}

func NewCreateUserBusiness(store CreateUserStore) *createUserBusiness {
	return &createUserBusiness{store: store}
}

func(c *createUserBusiness) CreateUser(newUser usermodel.NewUser, tokenSecret string) (string, error) {
	checkEmailExist := c.store.CheckEmailExist(newUser.Email)
	if checkEmailExist != nil {
		return "", common.NewErrorResponse(http.StatusBadRequest, checkEmailExist, 
			"Email đã tồn tại trong hệ thống", 
			"Email đã tồn tại trong hệ thống", 
			"Email đã tồn tại trong hệ thống")
	}

	hashedPassword, errHash := common.HashAndSalt([]byte(newUser.Password))
	if errHash != nil {
		return "", common.NewErrorResponse(http.StatusInternalServerError, errHash, 
			"Không thể đăng kí tài khoản", 
			"Không thể đăng kí tài khoản", 
			"Không thể đăng kí tài khoản")
	}

	var user usermodel.User
	user.Id = uuid.NewV4().String()
	user.Email = newUser.Email
	user.Password = hashedPassword

	userID, err := c.store.Create(user)
	if err != nil {
		return "", common.NewErrorResponse(http.StatusInternalServerError, err, 
			"Không thể đăng kí tài khoản", 
			"Không thể đăng kí tài khoản", 
			"Không thể đăng kí tài khoản")
	}

	tokenString, err := common.NewToken(userID, tokenSecret)
	if err != nil {
		return "", common.NewErrorResponse(http.StatusInternalServerError, err, 
			"Không thể tạo token", 
			"Không thể tạo token", 
			"Không thể tạo token")
	}

	return tokenString, nil
}