package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khunmostz/be-liongate-go/app/core/domain"
	"github.com/khunmostz/be-liongate-go/app/core/port"
)

type ShowRoundsController struct {
	svc port.ShowRoundsService
}

func NewShowRoundsController(svc port.ShowRoundsService) *ShowRoundsController {
	return &ShowRoundsController{
		svc: svc,
	}
}

func (src *ShowRoundsController) RegisterRoutes(router *gin.Engine) {
	showRounds := router.Group("/api/v1/show-rounds")
	{
		showRounds.GET("/", src.GetAllShowRounds)
		showRounds.POST("/", src.CreateShowRound)
		showRounds.POST("/:id", src.GetShowRoundById)
		showRounds.GET("/:id", src.GetShowRoundById)
		showRounds.PUT("/:id", src.UpdateShowRound)
		showRounds.DELETE("/:id", src.DeleteShowRound)
	}
}

// GetAllShowRounds godoc
// @Summary Get all show rounds
// @Description Get a list of all show rounds
// @Tags show-rounds
// @Accept json
// @Produce json
// @Success 200 {array} domain.ShowRounds
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /show-rounds [get]
func (src *ShowRoundsController) GetAllShowRounds(c *gin.Context) {
	showRounds, err := src.svc.GetAllShowRounds(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, showRounds)
}

// CreateShowRound godoc
// @Summary Create a new show round
// @Description Create a new show round with the provided information
// @Tags show-rounds
// @Accept json
// @Produce json
// @Param showRound body domain.ShowRounds true "Show Round information"
// @Success 201 {object} domain.ShowRounds
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /show-rounds [post]
func (src *ShowRoundsController) CreateShowRound(c *gin.Context) {
	var showRound domain.ShowRounds
	if err := c.ShouldBindJSON(&showRound); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := src.svc.CreateShowRound(c, &showRound)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}

// GetShowRoundById godoc
// @Summary Get a show round by ID
// @Description Get a show round's information by its ID
// @Tags show-rounds
// @Accept json
// @Produce json
// @Param id path string true "Show Round ID"
// @Success 200 {object} domain.ShowRounds
// @Failure 404 {object} map[string]interface{} "Show round not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /show-rounds/{id} [get]
// @Router /show-rounds/{id} [post]
func (src *ShowRoundsController) GetShowRoundById(c *gin.Context) {
	id := c.Param("id")
	showRound, err := src.svc.GetShowRoundById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if showRound == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Show round not found"})
		return
	}

	c.JSON(http.StatusOK, showRound)
}

// UpdateShowRound godoc
// @Summary Update a show round
// @Description Update a show round's information
// @Tags show-rounds
// @Accept json
// @Produce json
// @Param id path string true "Show Round ID"
// @Param showRound body domain.ShowRounds true "Updated Show Round information"
// @Success 200 {object} domain.ShowRounds
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 404 {object} map[string]interface{} "Show round not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /show-rounds/{id} [put]
func (src *ShowRoundsController) UpdateShowRound(c *gin.Context) {
	id := c.Param("id")

	// Check if show round exists
	_, err := src.svc.GetShowRoundById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedShowRound domain.ShowRounds
	if err := c.ShouldBindJSON(&updatedShowRound); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := src.svc.UpdateShowRound(c, id, &updatedShowRound)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// DeleteShowRound godoc
// @Summary Delete a show round
// @Description Delete a show round by its ID
// @Tags show-rounds
// @Accept json
// @Produce json
// @Param id path string true "Show Round ID"
// @Success 200 {object} map[string]interface{} "Success message"
// @Failure 404 {object} map[string]interface{} "Show round not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /show-rounds/{id} [delete]
func (src *ShowRoundsController) DeleteShowRound(c *gin.Context) {
	id := c.Param("id")

	err := src.svc.DeleteShowRound(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Show round deleted successfully"})
}
