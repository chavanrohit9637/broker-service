package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIHandler struct {
}

type APIResponse struct {
	Message string
}

func (api APIHandler) Broker(w http.ResponseWriter, r *http.Request) {

	fmt.Println("broker API called.. ")
	resp := APIResponse{
		Message: "Successful",
	}

	re, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(re)
}
