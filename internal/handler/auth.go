package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/abhiraj-ku/health_app/internal/auth"
	"github.com/abhiraj-ku/health_app/internal/model"
	"github.com/abhiraj-ku/health_app/internal/worker"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type AuthService interface {
	Authenticate(name string, password string) (*model.User, error)
	Register(user *model.User) (*model.User, error)
}

type AuthHandler struct {
	Service     AuthService
	RedisClient *redis.Client
}

func NewAuthHandler(s AuthService, redisClient *redis.Client) *AuthHandler {
	return &AuthHandler{
		Service:     s,
		RedisClient: redisClient,
	}
}

type RegisterRequest struct {
	Name      string    `json:"name" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Role      string    `json:"role" binding:"required"` // e.g., "doctor", "receptionist"+
	CreatedAt time.Time `json:"created_at"`
}

type Request struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("⚠️ Invalid register request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid field values"})
		return
	}

	user := &model.User{
		Name:      req.Name,
		Password:  req.Password,
		Role:      model.Role(req.Role),
		CreatedAt: time.Now(),
	}

	createdUser, err := h.Service.Register(user)
	log.Println(createdUser)
	if err != nil {
		log.Printf("❌ Failed to register user '%s': %v", req.Name, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register user"})
		return
	}

	// Enqueue welcome email task
	emailWorker := worker.NewEmailWorker(h.RedisClient)
	emailWorker.EnqueueEmail(createdUser)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "user registered successfully",
		"data":    createdUser,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req Request

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("⚠️ Invalid login request: %v", err)

		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid field values"})
		return
	}

	user, err := h.Service.Authenticate(req.Name, req.Password)
	if err != nil {
		log.Printf("❌ Login failed for user '%s': %v", user.Name, err)

		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication failed "})
		return
	}

	token, err := auth.GenerateToken(user.ID, string(user.Role))
	log.Printf("the token : %v", token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "failed to generate auth token "})
		return
	}

	// Enqueue the email task to Redis
	emailWorker := worker.NewEmailWorker(h.RedisClient) // Create new instance of email worker
	emailWorker.EnqueueEmail(user)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "user login sucess", "token": token, "data": user})
}
