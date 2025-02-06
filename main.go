package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fajarherdian22/saving-plan-api/controller"
	"github.com/fajarherdian22/saving-plan-api/db"
	"github.com/fajarherdian22/saving-plan-api/repository"
	"github.com/fajarherdian22/saving-plan-api/service"
	"github.com/fajarherdian22/saving-plan-api/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config :", err)
	}
	dbCon := db.ConDB(config.DBDriver, config.DBSource)
	repo := repository.New(dbCon)
	validate := validator.New()

	userService := service.NewUserService(repo)
	userController := controller.NewUserController(userService, validate)

	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"POST", "GET"},
		AllowHeaders:    []string{"Content-Type", "Origin"},
		ExposeHeaders:   []string{"Content-Length"},
		MaxAge:          12 * time.Hour,
	}))
	router.Use(gin.Recovery())
	r := router.Group("/api/")

	r.POST("/user/create", userController.CreateUser)
	r.GET("/user/email", userController.GetUser)
	err = router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
	fmt.Println("Service is Running")
}
