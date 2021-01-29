package controller

import (
	"food-delivery/domain"
	"food-delivery/common"

	"net/http"
	"strconv"

	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
)

func GetAllFood(provider common.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		result, err := noteHandler.NoteService.GetAll()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, result)
	}
	
}

func GetFoodById(provider common.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		idInt, _ := strconv.Atoi(id)

		result, err := noteHandler.NoteService.GetById(idInt)
		if err != nil {
			if err.Error() == "Không tìm thấy note" {
				c.JSON(http.StatusNotFound, domain.StatusMessage{
					Message: "Không tìm thấy note",
				})
				return
			}

			c.JSON(http.StatusInternalServerError, domain.StatusMessage{
				Message: "Lỗi server",
			})
			return
		}

		c.JSON(http.StatusOK, result)
	}
	
}
