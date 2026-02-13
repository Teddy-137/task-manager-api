package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/teddy-137/task_manager_api/internal/delivery"
	"github.com/teddy-137/task_manager_api/internal/domain"
	"github.com/teddy-137/task_manager_api/internal/repository"
	"github.com/teddy-137/task_manager_api/internal/service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database.")
	}

	db.AutoMigrate(&domain.User{}, &domain.Task{})

	router := gin.Default()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	delivery.NewUserHandler(router, userService)

	taskRepository := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepository)

	delivery.NewTaskHandler(router, taskService)

	log.Println("Starting Server on: 8080")
	router.Run()

}
