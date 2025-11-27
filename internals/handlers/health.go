package handlers

// import (
// 	"encoding/json"
// 	"net/http"
// )

// func (s *ApiServer) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
// 	response := struct {
// 		Status string `json:"status"`
// 	}{
// 		Status: "ok",
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(response)
// }
