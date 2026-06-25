package dto

type ErrorResponse struct {
	Code    string   `json:"code,omitempty"`
	Message string   `json:"message"`
	Details []string `json:"details,omitempty"`
}

func NewErrorResponse(code string, message string, details ...string) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: message,
		Details: details,
	}
}
