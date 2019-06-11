package utils

import (
	"encoding/json"
	"net/http"
)

// 返回统一格式, 
func Message(status bool, message string) (map[string]interface{}){
	return map[string]interface{}{
		"status":status,
		"message":message,
	}
}

// make map style to json style
func Respond(w http.ResponseWriter, data map[string] interface{}){
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}