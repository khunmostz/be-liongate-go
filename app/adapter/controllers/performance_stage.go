package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khunmostz/be-liongate-go/app/core/domain"
	"github.com/khunmostz/be-liongate-go/app/core/port"
)

type PerformanceStageController struct {
	svc port.PerformanceStageService
}

func NewPerformanceStageController(svc port.PerformanceStageService) *PerformanceStageController {
	return &PerformanceStageController{
		svc: svc,
	}
}

func (pc *PerformanceStageController) RegisterRoutes(router *gin.Engine) {
	stages := router.Group("/api/v1/stages")
	stages.GET("/", pc.GetStages)
	stages.POST("/", pc.CreateStage)
	stages.GET("/:id", pc.GetStageById)
	stages.PUT("/:id", pc.UpdateStage)
	stages.DELETE("/:id", pc.DeleteStage)
}

// GetStages godoc
// @Summary Get all performance stages
// @Description Get a list of all performance stages
// @Tags stages
// @Accept json
// @Produce json
// @Success 200 {array} domain.PerformanceStage
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /stages [get]
func (pc *PerformanceStageController) GetStages(c *gin.Context) {
	stages, err := pc.svc.GetStages(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stages)
}

// CreateStage godoc
// @Summary Create a new performance stage
// @Description Create a new performance stage with the provided information
// @Tags stages
// @Accept json
// @Produce json
// @Param stage body domain.PerformanceStage true "Performance Stage information"
// @Success 201 {object} domain.PerformanceStage
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /stages [post]
func (pc *PerformanceStageController) CreateStage(c *gin.Context) {
	var stage domain.PerformanceStage
	if err := c.ShouldBindJSON(&stage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := pc.svc.CreateStage(c, &stage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}

// GetStageById godoc
// @Summary Get a performance stage by ID
// @Description Get a performance stage's information by its ID
// @Tags stages
// @Accept json
// @Produce json
// @Param id path string true "Performance Stage ID"
// @Success 200 {object} domain.PerformanceStage
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /stages/{id} [get]
func (pc *PerformanceStageController) GetStageById(c *gin.Context) {
	id := c.Param("id")
	stage, err := pc.svc.GetStageById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stage)
}

// UpdateStage godoc
// @Summary Update a performance stage
// @Description Update a performance stage's information
// @Tags stages
// @Accept json
// @Produce json
// @Param id path string true "Performance Stage ID"
// @Param stage body domain.PerformanceStage true "Updated Performance Stage information"
// @Success 200 {object} domain.PerformanceStage
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /stages/{id} [put]
func (pc *PerformanceStageController) UpdateStage(c *gin.Context) {
	id := c.Param("id")
	// Check if stage exists
	_, err := pc.svc.GetStageById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedStage domain.PerformanceStage
	if err := c.ShouldBindJSON(&updatedStage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := pc.svc.UpdateStage(c, id, &updatedStage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// DeleteStage godoc
// @Summary Delete a performance stage
// @Description Delete a performance stage by its ID
// @Tags stages
// @Accept json
// @Produce json
// @Param id path string true "Performance Stage ID"
// @Success 200 {object} map[string]interface{} "Success message"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /stages/{id} [delete]
func (pc *PerformanceStageController) DeleteStage(c *gin.Context) {
	id := c.Param("id")
	err := pc.svc.DeleteStage(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Stage deleted successfully"})
}
