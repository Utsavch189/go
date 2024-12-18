package responses

import "time"

func ErrorResponse(err error, message string) map[string]interface{} {
	errorMessage := ""
	if err != nil {
		errorMessage = err.Error()
	}
	response := map[string]interface{}{
		"detail":    message,
		"error":     errorMessage,
		"timestamp": time.Now().Format(time.RFC3339),
	}
	return response
}
