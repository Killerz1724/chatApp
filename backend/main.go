package main

import (
	"chatApp/database"
	"chatApp/handler"
	"chatApp/middleware"
	"chatApp/repository"
	"chatApp/usecase"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := database.InitializeDatabase()

	if err != nil {
		// log.Fatalln(err)
		return
	}

	r := gin.New()
	r.ContextWithFallback = true
	r.Use(middleware.Logger())
	r.Use(middleware.ErrorMiddleware())

	ar := repository.NewAuthRepo(db)
	au := usecase.NewAuthUsecase(ar)
	ah := handler.NewAuthHandler(au)

	ur := repository.NewUserRepo(db)
	uu := usecase.NewUserUsecase(ur)
	uh := handler.NewUserHandler(uu)

	cr := repository.NewConvRepository(db)
	cu := usecase.NewConvUsecase(cr)
	ch := handler.NewConvHandler(cu)

	{
		auth := r.Group("/api/auth")
		auth.POST("/login", ah.HandlerLogin)
		auth.POST("/register", ah.HandlerAuthRegister)
	}
	{
		user := r.Group("/api/user")
		user.Use(middleware.AuthMiddleware())
		user.GET("/me", uh.HandlerGetUserProfile)
		user.GET("/friends", uh.UsecaseGetUserFriends)
		user.GET("/friends/:id", uh.HandlerGetUserDetailFriend)
	}
	{
		conversations := r.Group("/api/conversations")
		conversations.Use(middleware.AuthMiddleware())
		conversations.GET("", ch.HandlerGetListConv)
	}


	srv := &http.Server{
		Addr:    ":" + os.Getenv("SERVER_PORT"),
		Handler: r.Handler(),
	}

	log.Println("Listening on port: " + os.Getenv("SERVER_PORT"))

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatalln(err)
		return
	}

	defer db.Close()

}