package responses

import (
	"time"

	"github.com/Utsavch189/api_mysql/internal/models/requests"
)

func CreateProductRepsonse(product *requests.Product, detail string) map[string]interface{} {
	response := map[string]interface{}{
		"product":   product,
		"timestamp": time.Now().Format(time.RFC3339),
		"detail":    detail,
	}
	return response
}

func UpdateProductRepsonse(product *requests.Product, detail string) map[string]interface{} {
	response := map[string]interface{}{
		"product":   product,
		"timestamp": time.Now().Format(time.RFC3339),
		"detail":    detail,
	}
	return response
}

func DeleteProductRepsonse(message string) map[string]interface{} {
	response := map[string]interface{}{
		"timestamp": time.Now().Format(time.RFC3339),
		"detail":    message,
	}
	return response
}

func GetAllProductResponse(products []requests.Product) map[string]interface{} {
	response := map[string]interface{}{
		"products":  products,
		"timestamp": time.Now().Format(time.RFC3339),
	}
	return response
}

func GetAProductResponse(product requests.Product) map[string]interface{} {
	response := map[string]interface{}{
		"product":   product,
		"timestamp": time.Now().Format(time.RFC3339),
	}
	return response
}
