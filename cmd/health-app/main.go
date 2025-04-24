package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/abhiraj-ku/health_app/config"
	"github.com/abhiraj-ku/health_app/internal/auth"
	"github.com/abhiraj-ku/health_app/internal/handler"
	postresdb "github.com/abhiraj-ku/health_app/internal/repository/db"
	"github.com/abhiraj-ku/health_app/internal/service"
	"github.com/abhiraj-ku/health_app/internal/worker"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func main() {
	config.LoadConfig()
	db, err := sql.Open("postgres", config.AppConfig.PostgresURI)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	// Connect to Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	//  Initialize user's repository,services and handlers
	userRepo := postresdb.NewUserRepo(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService, redisClient)

	// Initialize patient repository,services and handlers

	patientRepo := postresdb.NewPatientRepo(db)
	patientService := service.NewPatientService(patientRepo)
	patientHandler := handler.NewPatientHandler(patientService)

	// Initialize the email worker
	emailWorker := worker.NewEmailWorker(redisClient)
	go emailWorker.ProcessEmailQueue() // Start background worker for email queue

	router := gin.Default()
	router.POST("/login", authHandler.Login)

	// Patient routes (RBAC controlled)
	patientHandler.RegisterRoutes(router, auth.JWTMiddleware(), auth.RequireRole)

	// Check health of server
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.Run(config.AppConfig.ServerPort)
}
