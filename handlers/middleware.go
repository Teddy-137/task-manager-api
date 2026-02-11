package handlers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func authMiddleware(ctx *gin.Context) {
	apiKey := ctx.GetHeader("X-API-KEY")
	validKey := os.Getenv("API_KEY")
	if apiKey != validKey {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		ctx.Abort()
		return
	}

	ctx.Next()
}

func logRequest(ctx *gin.Context) {
	t := time.Now()
	ctx.Next()

	latency := time.Since(t)
	log.Printf("--Path: %s | Status: %d | Latency: %v", ctx.Request.URL.Path, ctx.Writer.Status(), latency)
}

func jsonRecovery(ctx *gin.Context) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("error: ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error."})
			ctx.Abort()
			return
		}
	}()

	ctx.Next()
}
