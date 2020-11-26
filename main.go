package main

import (
	"log"
	"os"

	"creating-simple-api/model"
	questionHandler "creating-simple-api/question/handler"
	questionRepo "creating-simple-api/question/repo"
	questionUsecase "creating-simple-api/question/usecase"

	userRepo "creating-simple-api/user/repo"
	userUsecase "creating-simple-api/user/usecase"
	userHandler "creating-simple-api/user/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload"
)


func main() {
	db, err := gorm.Open("mysql", os.Getenv("MYSQL"))
	if err != nil {
		log.Fatal(err)
	}

	db.Debug().AutoMigrate(
		&model.Question{},
		&model.User{},
	)

	router := gin.Default()

	questionRepo := questionRepo.CreateQuestionRepo(db)
	questionUsecase := questionUsecase.CreateQuestionUsecase(questionRepo)
	userRepo := userRepo.CreateUserRepo(db)
	userUsecase := userUsecase.CreateUserUsecase(userRepo)

	questionHandler.CreateQuestionHandler(router, questionUsecase, userUsecase)
	userHandler.CreateUserHandler(router, userUsecase)


	err = router.Run(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
}