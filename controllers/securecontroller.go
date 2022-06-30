package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(context *gin.Context) {
	fmt.Println("Securely Access to your Application")
	context.JSON(http.StatusOK, gin.H{"message": "pong"})
}
