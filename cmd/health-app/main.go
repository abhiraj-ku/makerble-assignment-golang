package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/abhiraj-ku/health_app/config"
	"github.com/abhiraj-ku/health_app/internal/auth"
	"github.com/abhiraj-ku/health_app/internal/handler"
	postresdb "github.com/abhiraj-ku/health_app/internal/repository/db"
	"github.com/abhiraj-ku/health_app/internal/service"
	"github.com/abhiraj-ku/health_app/internal/worker"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
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

	// Start HTTP server in a separate goroutine
	server := &http.Server{
		Addr:    config.AppConfig.ServerPort,
		Handler: router,
	}

	// Channel to listen for termination signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Run the server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
		log.Printf("server is up and running on the port: %v", config.AppConfig.ServerPort)
	}()

	// Wait for an interrupt signal
	<-stop
	log.Println("Shutting down gracefully...")

	// Set a timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to shut down the server gracefully
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server exited properly")
}
