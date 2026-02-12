package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teddy-137/task_manager_api/internal/domain"
	"github.com/teddy-137/task_manager_api/internal/middleware"
)

type UserHandler struct {
	Service domain.UserService
}

func NewUserHandler(r *gin.Engine, s domain.UserService) {
	handler := &UserHandler{
		Service: s,
	}

	r.POST("/login", handler.Login)
	r.POST("/register", handler.CreateUser)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/users", handler.GetAllUsers)
	}
}
func (h *UserHandler) GetAllUsers(ctx *gin.Context) {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var user domain.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (h *UserHandler) Login(ctx *gin.Context) {
	var input struct {
		username string
		password string
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.Service.Login(input.username, input.password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})

}
