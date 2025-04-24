package handler

import (
	"net/http"

	"github.com/abhiraj-ku/health_app/internal/auth"
	"github.com/abhiraj-ku/health_app/internal/model"
	"github.com/abhiraj-ku/health_app/internal/worker"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type AuthService interface {
	Authenticate(username string, password string) (*model.User, error)
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

type Request struct {
	Username string `json:"username"`
	Password string `'json:"username"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req Request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid field values"})
		return
	}

	user, err := h.Service.Authenticate(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication failed "})
		return
	}

	token, err := auth.GenerateToken(user.ID, string(user.Role))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "failed to generate auth token "})
		return
	}

	// Enqueue the email task to Redis
	emailWorker := worker.NewEmailWorker(h.RedisClient) // Create new instance of email worker
	emailWorker.EnqueueEmail(user)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "user login sucess", "token": token, "data": user})
}
