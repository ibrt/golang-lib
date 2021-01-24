package errors

// ErrorResponse describes an error response.
type ErrorResponse struct {
	ErrorID    ID                     `json:"errorId,omitempty"`
	StatusCode StatusCode             `json:"statusCode,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
	Message    string                 `json:"message,omitempty"`
	StackTrace []string               `json:"stackTrace,omitempty"`
}

// ToResponse converts an error to ErrorResponse.
func ToResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		ErrorID:    GetID(err),
		StatusCode: GetStatusCode(err),
		Metadata:   GetMetadata(err),
		Message:    err.Error(),
		StackTrace: FormatStackTrace(GetCallers(err)),
	}
}
