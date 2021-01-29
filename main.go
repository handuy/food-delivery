package main

import (
	"food-delivery/config"
	// "food-delivery/component/appctx"

	userController "food-delivery/module/user/controller"
	userRepository "food-delivery/module/user/repository/mysql"
	userServ "food-delivery/user/service"

	noteController "food-delivery/note/controller"
	noteRepository "food-delivery/note/repository/mysql"
	noteServ "food-delivery/note/service"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	dbConn, tokenSecret, err := config.SetUpDbAndSecret(".")
	if err != nil {
		log.Fatal(err)
	}

	// appCtx := appctx.New(dbConn)

	r := gin.Default()
	groupRoutes := r.Group("/")

	userRepo := userRepository.NewUserRepository(dbConn)
	userService := userServ.NewUserService(userRepo)

	noteRepo := noteRepository.NewNoteRepo(dbConn)
	noteService := noteServ.NewNoteService(noteRepo)

	userController.NewUserHandler(groupRoutes, userService, tokenSecret)
	noteController.NewNoteHandler(groupRoutes, noteService, userService, tokenSecret)

	r.Run("localhost:8181")
}
