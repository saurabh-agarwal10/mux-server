package utils

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

// Response returns an http response with the payload provided
func Response(w http.ResponseWriter, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		zap.L().Fatal("Error in response")
		return
	}
}
