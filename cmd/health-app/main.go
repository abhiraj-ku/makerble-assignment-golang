package main

import (
	"github.com/abhiraj-ku/health_app/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	router := gin.Default()
}
