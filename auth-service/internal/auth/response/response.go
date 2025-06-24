package response

type ServiceResponse[T any] struct {
	Success bool          `json:"success"`
	Data    *T            `json:"data,omitempty"`
	Error   *ServiceError `json:"error,omitempty"`
}

type ServiceError struct {
	Code    string `json:"code"`    // "VALIDATION_ERROR"
	Message string `json:"message"` // "user_id must be a UUID"
	Field   string `json:"field,omitempty"`
}

// UI response
type APIResponse[T any] struct {
	Success bool   `json:"success"`
	Data    *T     `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}
