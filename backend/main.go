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

	{
		auth := r.Group("/api/auth")
		auth.POST("/login", ah.HandlerLogin)
		auth.POST("/register", ah.HandlerAuthRegister)
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