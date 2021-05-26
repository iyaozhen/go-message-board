package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/sessions"
	"mydemo/config"
	"net/http"
	"strconv"
)

func BaseResponse(w http.ResponseWriter, dataCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responseData := make(map[string]interface{})
	responseData["data_code"] = strconv.Itoa(dataCode)
	responseData["data"] = data

	err := json.NewEncoder(w).Encode(responseData)
	if err != nil {
		fmt.Fprintf(w, "json error")
	}
}

func SessionStart() *sessions.FilesystemStore {
	conf := config.Config().Security
	store := sessions.NewFilesystemStore("", []byte(conf["session_key"].(string)))

	return store
}
