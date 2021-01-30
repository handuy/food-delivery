package foodtransport

import (
	"food-delivery/common"
	"food-delivery/module/food/foodbusiness"
	"food-delivery/module/food/foodstorage"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllFood(provider common.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := provider.GetMainDBConnection()
		sqlStore := foodstorage.NewFoodStorage(db)
		foodBiz := foodbusiness.NewGetFoodBusiness(sqlStore)

		result, err := foodBiz.GetAll()
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

		db := provider.GetMainDBConnection()
		sqlStore := foodstorage.NewFoodStorage(db)
		foodBiz := foodbusiness.NewGetFoodBusiness(sqlStore)

		result, err := foodBiz.GetById(idInt)
		if err != nil {
			if err.Error() == "Không tìm thấy note" {
				c.JSON(http.StatusNotFound, common.NewErrorResponse(
					http.StatusNotFound, err,
					"Không tìm thấy note", "Không tìm thấy note", "Không tìm thấy note",
				))
				return
			}

			c.JSON(http.StatusInternalServerError, common.NewErrorResponse(
				http.StatusNotFound, err,
				"Không tìm thấy note", "Không tìm thấy note", "Không tìm thấy note",
			))
			return
		}

		c.JSON(http.StatusOK, result)
	}

}
