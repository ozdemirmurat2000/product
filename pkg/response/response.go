package response

type ErrorResponseModel struct {
	Success bool   `json:"success" example:"false"`
	Error   string `json:"error,omitempty" example:"Error message"`
}
type SuccessResponseModel struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message,omitempty" example:"Operation successful"`
	Data    any    `json:"data,omitempty"`
}

func SuccessResponse(message string, data any) SuccessResponseModel {
	return SuccessResponseModel{
		Success: true,
		Data:    data,
		Message: message,
	}
}

func ErrorResponse(message string) ErrorResponseModel {
	return ErrorResponseModel{
		Success: false,
		Error:   message,
	}
}
