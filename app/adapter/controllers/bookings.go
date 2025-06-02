package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khunmostz/be-liongate-go/app/core/domain"
	"github.com/khunmostz/be-liongate-go/app/core/port"
)

type BookingsController struct {
	svc port.BookingsService
}

func NewBookingsController(svc port.BookingsService) *BookingsController {
	return &BookingsController{
		svc: svc,
	}
}

func (bc *BookingsController) RegisterRoutes(router *gin.Engine) {
	bookings := router.Group("/api/v1/bookings")
	{
		bookings.POST("", bc.CreateBooking)
		bookings.GET("/:id", bc.GetBookingById)
		bookings.GET("/user/:userId", bc.GetBookingsByUserId)
		bookings.GET("/round/:roundId", bc.GetBookingsByRoundId)
		bookings.PUT("/:id", bc.UpdateBooking)
		bookings.DELETE("/:id", bc.DeleteBooking)
	}
}

// CreateBooking godoc
// @Summary Create a new booking
// @Description Create a new booking with the provided information
// @Tags bookings
// @Accept json
// @Produce json
// @Param booking body domain.Bookings true "Booking information"
// @Success 201 {object} domain.Bookings
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /bookings [post]
func (bc *BookingsController) CreateBooking(c *gin.Context) {
	var booking domain.Bookings
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := bc.svc.CreateBooking(c.Request.Context(), &booking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

// GetBookingById godoc
// @Summary Get a booking by ID
// @Description Get a booking's information by its ID
// @Tags bookings
// @Accept json
// @Produce json
// @Param id path string true "Booking ID"
// @Success 200 {object} domain.Bookings
// @Failure 404 {object} map[string]interface{} "Booking not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /bookings/{id} [get]
func (bc *BookingsController) GetBookingById(c *gin.Context) {
	id := c.Param("id")
	booking, err := bc.svc.GetBookingById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if booking == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	c.JSON(http.StatusOK, booking)
}

// GetBookingsByUserId godoc
// @Summary Get bookings by user ID
// @Description Get all bookings for a specific user
// @Tags bookings
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {array} domain.Bookings
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /bookings/user/{userId} [get]
func (bc *BookingsController) GetBookingsByUserId(c *gin.Context) {
	userId := c.Param("userId")
	bookings, err := bc.svc.GetBookingsByUserId(c.Request.Context(), userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bookings)
}

// GetBookingsByRoundId godoc
// @Summary Get bookings by round ID
// @Description Get all bookings for a specific show round
// @Tags bookings
// @Accept json
// @Produce json
// @Param roundId path string true "Round ID"
// @Success 200 {array} domain.Bookings
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /bookings/round/{roundId} [get]
func (bc *BookingsController) GetBookingsByRoundId(c *gin.Context) {
	roundId := c.Param("roundId")
	bookings, err := bc.svc.GetBookingsByRoundId(c.Request.Context(), roundId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bookings)
}

// UpdateBooking godoc
// @Summary Update a booking
// @Description Update a booking's information
// @Tags bookings
// @Accept json
// @Produce json
// @Param id path string true "Booking ID"
// @Param booking body domain.Bookings true "Updated Booking information"
// @Success 200 {object} domain.Bookings
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 404 {object} map[string]interface{} "Booking not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /bookings/{id} [put]
func (bc *BookingsController) UpdateBooking(c *gin.Context) {
	id := c.Param("id")

	// Check if booking exists
	_, err := bc.svc.GetBookingById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedBooking domain.Bookings
	if err := c.ShouldBindJSON(&updatedBooking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := bc.svc.UpdateBooking(c.Request.Context(), id, &updatedBooking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// DeleteBooking godoc
// @Summary Delete a booking
// @Description Delete a booking by its ID
// @Tags bookings
// @Accept json
// @Produce json
// @Param id path string true "Booking ID"
// @Success 200 {object} map[string]interface{} "Success message"
// @Failure 404 {object} map[string]interface{} "Booking not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /bookings/{id} [delete]
func (bc *BookingsController) DeleteBooking(c *gin.Context) {
	id := c.Param("id")

	err := bc.svc.DeleteBooking(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking deleted successfully"})
}
