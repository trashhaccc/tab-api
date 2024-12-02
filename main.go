package main

import (
	"github.com/HIUNCY/rest-api-go/handler"
	"github.com/HIUNCY/rest-api-go/model"
	"github.com/HIUNCY/rest-api-go/repository"
	"github.com/HIUNCY/rest-api-go/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:CRNkojvoRMyfovNjlCKNebPnSrEhOxvS@tcp(junction.proxy.rlwy.net:26272)/railway?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.User{}, &model.Transaction{})

	// USER
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	// TRANSACTION
	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepository)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	r := gin.Default()
	r.Use(cors.Default())
	user := r.Group("/user")
	{
		user.GET("/list", userHandler.GetUserList)
		// user.GET("/balance/:nik", userHandler.GetUserList)
		user.POST("/login", userHandler.Login)
		user.POST("/register", userHandler.Register)
		user.PUT("/update", userHandler.Update)
		user.DELETE("/delete", userHandler.Delete)
	}
	transaction := r.Group("/transaction")
	{
		transaction.POST("/create", transactionHandler.CreateTransaction)
		transaction.GET("/history/:nik", transactionHandler.HistoryTransaction)
	}
	r.Run()
}
