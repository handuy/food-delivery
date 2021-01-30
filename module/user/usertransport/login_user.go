package usertransport

import (
	"food-delivery/common"
	"food-delivery/module/user/userbusiness"
	"food-delivery/module/user/usermodel"
	"food-delivery/module/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginUser(provider common.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var newUser usermodel.NewUser
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, common.NewErrorResponse(http.StatusBadRequest, err,
				"Yêu cầu không hợp lệ",
				"Yêu cầu không hợp lệ",
				"Yêu cầu không hợp lệ"))
			return
		}

		db := provider.GetMainDBConnection()
		sqlStore := userstorage.NewSQLStore(db)
		loginUserBusiness := userbusiness.NewLoginBusiness(sqlStore)
		tokenSecret := provider.GetTokenSecret()

		token, err := loginUserBusiness.LogIn(newUser, tokenSecret)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewErrorResponse(http.StatusBadRequest, err,
				"Không thể đăng kí tài khoản",
				"Không thể đăng kí tài khoản",
				"Không thể đăng kí tài khoản"))
			return
		}

		c.JSON(200, gin.H{"token": token})
	}
}
