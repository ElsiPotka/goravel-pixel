package resources

type ApiResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func NewSuccessResponse(message string, data any) *ApiResponse {
	return &ApiResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}

func NewErrorResponse(message string, data any) *ApiResponse {
	return &ApiResponse{
		Status:  "error",
		Message: message,
		Data:    data,
	}
}

func NewValidationErrorResponse(errors any) *ApiResponse {
	return &ApiResponse{
		Status:  "error",
		Message: "Validation failed",
		Data:    errors,
	}
}
