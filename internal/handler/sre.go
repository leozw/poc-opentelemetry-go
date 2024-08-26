package handler

import (
	"modulo/pkg/response"
	"net/http"
)

func SREHandler(w http.ResponseWriter, r *http.Request) {

	response.JSON(w, http.StatusOK, map[string]string{"message": "Welcome to the SRE page!"})
}
