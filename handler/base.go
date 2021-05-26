package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func BaseResponse(w http.ResponseWriter, dataCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responseData := make(map[string]interface{})
	responseData["dataCode"] = strconv.Itoa(dataCode)
	responseData["data"] = data

	err := json.NewEncoder(w).Encode(responseData)
	if err != nil {
		fmt.Fprintf(w, "json error")
	}
}
