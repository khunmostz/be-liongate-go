package domain

// ErrorResponse represents a standard error response structure for API endpoints
type ErrorResponse struct {
	Error string `json:"error" example:"An error occurred"`
}

// SuccessMessage represents a standard success message response
type SuccessMessage struct {
	Message string `json:"message" example:"Operation completed successfully"`
}
