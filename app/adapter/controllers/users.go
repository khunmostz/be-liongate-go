package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khunmostz/be-liongate-go/app/core/domain"
	"github.com/khunmostz/be-liongate-go/app/core/port"
)

type UsersController struct {
	svc port.UsersService
}

func NewUsersController(svc port.UsersService) *UsersController {
	return &UsersController{
		svc: svc,
	}
}

func (uc *UsersController) RegisterRoutes(router *gin.Engine) {
	users := router.Group("/api/v1/users")
	{
		users.POST("/register", uc.Register)
		users.GET("/:id", uc.GetUserById)
		users.GET("/role/:role", uc.GetUsersByRole)
		users.PUT("/:id", uc.UpdateUser)
		users.DELETE("/:id", uc.DeleteUser)
	}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param user body domain.Users true "User information"
// @Success 201 {object} domain.Users
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /users/register [post]
func (uc *UsersController) Register(c *gin.Context) {
	var user domain.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := uc.svc.Register(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

// GetUserById godoc
// @Summary Get a user by ID
// @Description Get a user's information by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} domain.Users
// @Failure 404 {object} map[string]interface{} "User not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /users/{id} [get]
func (uc *UsersController) GetUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := uc.svc.GetUserById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUsersByRole godoc
// @Summary Get users by role
// @Description Get all users with a specific role
// @Tags users
// @Accept json
// @Produce json
// @Param role path string true "User role"
// @Success 200 {array} domain.Users
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /users/role/{role} [get]
func (uc *UsersController) GetUsersByRole(c *gin.Context) {
	role := c.Param("role")
	users, err := uc.svc.GetUsersByRole(c.Request.Context(), role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update a user's information
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body domain.Users true "Updated User information"
// @Success 200 {object} domain.Users
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 404 {object} map[string]interface{} "User not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /users/{id} [put]
func (uc *UsersController) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	// Check if user exists
	_, err := uc.svc.GetUserById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedUser domain.Users
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := uc.svc.UpdateUser(c.Request.Context(), id, &updatedUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]interface{} "Success message"
// @Failure 404 {object} map[string]interface{} "User not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /users/{id} [delete]
func (uc *UsersController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	err := uc.svc.DeleteUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
