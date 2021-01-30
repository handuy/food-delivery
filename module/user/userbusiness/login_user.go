package userbusiness

import (
	"food-delivery/common"
	"food-delivery/module/user/usermodel"
	"net/http"
)

type LoginStore interface {
	GetUser(email string) (usermodel.User, error)
}

type loginBusiness struct {
	store LoginStore
}

func NewLoginBusiness(store LoginStore) *loginBusiness {
	return &loginBusiness{store: store}
}

func (l *loginBusiness) LogIn(loginUser usermodel.NewUser, tokenSecret string) (string, error) {
	userInfo, err := l.store.GetUser(loginUser.Email)
	if err != nil {
		return "", common.NewErrorResponse(http.StatusUnauthorized, err,
			"Sai thông tin đăng nhập",
			"Sai thông tin đăng nhập",
			"Sai thông tin đăng nhập")
	}

	comparePass := common.CheckPassword([]byte(userInfo.Password), []byte(loginUser.Password))
	if comparePass != nil {
		return "", common.NewErrorResponse(http.StatusUnauthorized, err,
			"Sai thông tin đăng nhập",
			"Sai thông tin đăng nhập",
			"Sai thông tin đăng nhập")
	}

	tokenString, err := common.NewToken(userInfo.Id, tokenSecret)
	if err != nil {
		return "", common.NewErrorResponse(http.StatusInternalServerError, err,
			"Không thể tạo token",
			"Không thể tạo token",
			"Không thể tạo token")
	}

	return tokenString, nil
}
