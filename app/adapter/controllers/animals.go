package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khunmostz/be-liongate-go/app/core/domain"
	"github.com/khunmostz/be-liongate-go/app/core/port"
)

type AnimalsController struct {
	svc          port.AnimalsService
	showRoundSvc port.ShowRoundsService
}

func NewAnimalsController(svc port.AnimalsService, showRoundSvc port.ShowRoundsService) *AnimalsController {
	return &AnimalsController{
		svc:          svc,
		showRoundSvc: showRoundSvc,
	}
}

func (ac *AnimalsController) RegisterRoutes(router *gin.Engine) {
	animals := router.Group("/api/v1/animals")
	animals.GET("/", ac.GetAnimals)
	animals.POST("/", ac.CreateAnimal)
	animals.GET("/:id", ac.GetAnimalById)
	animals.PUT("/:id", ac.UpdateAnimal)
	animals.DELETE("/:id", ac.DeleteAnimal)
	animals.POST("/:id/perform-show/:roundId", ac.PerformShowRound)
}

// GetAnimals godoc
// @Summary Get all animals
// @Description Get a list of all animals
// @Tags animals
// @Accept json
// @Produce json
// @Success 200 {array} domain.Animals
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /animals [get]
func (ac *AnimalsController) GetAnimals(c *gin.Context) {
	animals, err := ac.svc.GetAnimals(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, animals)
}

// CreateAnimal godoc
// @Summary Create a new animal
// @Description Create a new animal with the provided information
// @Tags animals
// @Accept json
// @Produce json
// @Param animal body domain.Animals true "Animal information"
// @Success 201 {object} domain.Animals
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /animals [post]
func (ac *AnimalsController) CreateAnimal(c *gin.Context) {
	var animal domain.Animals
	if err := c.ShouldBindJSON(&animal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := ac.svc.CreateAnimal(c, &animal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}

// GetAnimalById godoc
// @Summary Get an animal by ID
// @Description Get an animal's information by its ID
// @Tags animals
// @Accept json
// @Produce json
// @Param id path string true "Animal ID"
// @Success 200 {object} domain.Animals
// @Failure 404 {object} map[string]interface{} "Animal not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /animals/{id} [get]
func (ac *AnimalsController) GetAnimalById(c *gin.Context) {
	id := c.Param("id")
	animal, err := ac.svc.GetAnimalById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if animal == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Animal not found"})
		return
	}

	c.JSON(http.StatusOK, animal)
}

// UpdateAnimal godoc
// @Summary Update an animal
// @Description Update an animal's information
// @Tags animals
// @Accept json
// @Produce json
// @Param id path string true "Animal ID"
// @Param animal body domain.Animals true "Updated Animal information"
// @Success 200 {object} domain.Animals
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 404 {object} map[string]interface{} "Animal not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /animals/{id} [put]
func (ac *AnimalsController) UpdateAnimal(c *gin.Context) {
	id := c.Param("id")

	// Check if animal exists
	_, err := ac.svc.GetAnimalById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedAnimal domain.Animals
	if err := c.ShouldBindJSON(&updatedAnimal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := ac.svc.UpdateAnimal(c, id, &updatedAnimal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// DeleteAnimal godoc
// @Summary Delete an animal
// @Description Delete an animal by its ID
// @Tags animals
// @Accept json
// @Produce json
// @Param id path string true "Animal ID"
// @Success 200 {object} map[string]interface{} "Success message"
// @Failure 404 {object} map[string]interface{} "Animal not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /animals/{id} [delete]
func (ac *AnimalsController) DeleteAnimal(c *gin.Context) {
	id := c.Param("id")

	err := ac.svc.DeleteAnimal(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Animal deleted successfully"})
}

// PerformShowRound godoc
// @Summary Animal performs a show round
// @Description Record an animal performing a specific show round
// @Tags animals
// @Accept json
// @Produce json
// @Param id path string true "Animal ID"
// @Param roundId path string true "Show Round ID"
// @Success 200 {object} domain.ShowRounds
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 404 {object} map[string]interface{} "Animal or show round not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /animals/{id}/perform-show/{roundId} [post]
func (ac *AnimalsController) PerformShowRound(c *gin.Context) {
	animalId := c.Param("id")
	roundId := c.Param("roundId")

	// Verify animal exists
	animal, err := ac.svc.GetAnimalById(c, animalId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Animal not found"})
		return
	}

	// Verify show round exists
	showRound, err := ac.showRoundSvc.GetShowRoundById(c.Request.Context(), roundId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Show round not found"})
		return
	}

	// Verify this animal is assigned to this show round
	if showRound.AnimalId != animalId {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This animal is not assigned to this show round"})
		return
	}

	// Return the show round details with animal information
	response := gin.H{
		"show_round": showRound,
		"animal":     animal,
		"status":     "performing",
	}

	c.JSON(http.StatusOK, response)
}
