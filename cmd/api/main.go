package main

import (
	"log"

	"github.com/gin-gonic/gin"
	bookscontroller "github.com/marcelodiasdev/library-mvc/internal/books/controllers"
	loancontroller "github.com/marcelodiasdev/library-mvc/internal/loans/controllers"
	userscontroller "github.com/marcelodiasdev/library-mvc/internal/users/controllers"
)

func main() {
	router := gin.Default()
	usersController := userscontroller.NewUserController()
	booksController := bookscontroller.NewUserController()
	loansController := loancontroller.NewUserController()

	usersController.RegisterRoutes(router)
	booksController.RegisterRoutes(router)
	loansController.RegisterRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
