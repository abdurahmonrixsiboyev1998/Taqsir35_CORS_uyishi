package handlers

import (
	"context"
	"cors/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrigins(c *gin.Context) {
	origins, err := config.RedisClient.SMembers(context.Background(), "allowed_origins").Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch origins"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"origins": origins})
}

func AddOrigin(c *gin.Context) {
	var input struct {
		Origin string `json:"origin" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := config.RedisClient.SAdd(context.Background(), "allowed_origins", input.Origin).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not add origin"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Origin added successfully"})
}

func DeleteOrigin(c *gin.Context) {
	origin := c.Param("origin")
	err := config.RedisClient.SRem(context.Background(), "allowed_origins", origin).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete origin"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Origin deleted successfully"})
}
