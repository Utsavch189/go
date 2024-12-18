package responses

import (
	"time"

	"github.com/Utsavch189/api_mysql/internal/models/requests"
)

func CreateUserResponse(user *requests.User, detail string) map[string]interface{} {
	response := map[string]interface{}{
		"user":      user,
		"timestamp": time.Now().Format(time.RFC3339),
		"detail":    detail,
	}
	return response
}

func LoginResponse(token string, detail string) map[string]interface{} {
	response := map[string]interface{}{
		"jwt":       token,
		"timestamp": time.Now().Format(time.RFC3339),
		"detail":    detail,
	}
	return response
}
