package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khunmostz/be-liongate-go/app/core/domain"
	"github.com/khunmostz/be-liongate-go/app/core/port"
)

type AuthController struct {
	svc port.AuthService
}

func NewAuthController(svc port.AuthService) *AuthController {
	return &AuthController{svc: svc}
}

func (ac *AuthController) RegisterRoutes(router *gin.Engine) {
	router.POST("/api/v1/auth/login", ac.Login)
	router.POST("/api/v1/auth/register", ac.Register)
	router.POST("/api/v1/auth/refresh-token", ac.RefreshToken)
}

// Login godoc
// @Summary      User login
// @Description  Authenticate user with username and password
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        request body domain.LoginRequest true "Login credentials"
// @Success      200 {object} domain.AuthResponse "Successful login"
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router       /auth/login [post]
func (ac *AuthController) Login(c *gin.Context) {
	var req domain.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authResponse, err := ac.svc.Login(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, authResponse)
}

// Register godoc
// @Summary      User registration
// @Description  Register a new user account
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        request body domain.RegisterRequest true "Registration details"
// @Success      200 {object} domain.AuthResponse "Successful registration"
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router       /auth/register [post]
func (ac *AuthController) Register(c *gin.Context) {
	var req domain.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authResponse, err := ac.svc.Register(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, authResponse)
}

// RefreshToken godoc
// @Summary      Refresh access token
// @Description  Generate new access token using refresh token
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        request body domain.RefreshTokenRequest true "Refresh token"
// @Success      200 {object} domain.TokenPair "New token pair"
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router       /auth/refresh-token [post]
func (ac *AuthController) RefreshToken(c *gin.Context) {
	var req domain.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenPair, err := ac.svc.RefreshToken(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, tokenPair)
}
