package handler

import (
	"net/http"

	"github.com/abhiraj-ku/health_app/internal/auth"
	"github.com/abhiraj-ku/health_app/internal/model"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authenticate(username string, password string) (*model.User, error)
}

type AuthHandler struct {
	Service AuthService
}

func NewAuthHandler(s AuthService) *AuthHandler {
	return &AuthHandler{
		Service: s,
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

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "user login sucess", "token": token, "data": user})
}
