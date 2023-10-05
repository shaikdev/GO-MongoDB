package helper

import (
	"encoding/json"
	"net/http"
)

func ResponseErrorSender(w http.ResponseWriter, message string, status string, code int) {
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"message":    message,
		"status":     status,
		"statusCode": code,
	})
}

func ResponseSuccessSender(w http.ResponseWriter, message string, status string, code int, data interface{}) {
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"message":    message,
		"status":     status,
		"statusCode": code,
		"data":       data,
	})
}

func ResponseSuccessSenderWithoutData(w http.ResponseWriter, message string, status string, code int) {
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"message":    message,
		"status":     status,
		"statusCode": code,
	})
}

func ResponseSuccessSenderWithCount(w http.ResponseWriter, message string, status string, code int, data interface{}, count int) {
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"message":    message,
		"status":     status,
		"statusCode": code,
		"data":       data,
		"count":      count,
	})
}

func Header(w http.ResponseWriter, methods string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", methods)
}
