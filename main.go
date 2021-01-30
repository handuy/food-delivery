package main

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	"food-delivery/router"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	dbConn, tokenSecret, err := common.SetUpDbAndSecret(".")
	if err != nil {
		log.Fatal(err)
	}

	appContext := appctx.New(dbConn, tokenSecret)

	r := gin.Default()
	router.SetupRouter(r, appContext)

	r.Run("localhost:8181")
}
