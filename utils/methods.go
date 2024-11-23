package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webhookapi/models"
)

func WriteResponse(w http.ResponseWriter, res models.ApiResponse) {
	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(msg)
}

func ReturnError(msg string) models.ApiResponse {
	return models.ApiResponse{Code: -1, Message: msg, Data: nil}
}
