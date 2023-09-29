package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/victormagalhaess/go-api-template/pkg/api/services"
	"github.com/victormagalhaess/go-api-template/pkg/api/status"
	"github.com/victormagalhaess/go-api-template/pkg/model"
)

// @Description Tracker endpoint
// @Summary Endpoint to start the mining
// @Produce  text/plain
// @Router /api/v1/tracker [GET]
// @Success 200
// @Tags Tracker
func Tracker(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		status.BadRequest(w, err)
		return
	}

	request := &model.Request{}
	err = json.Unmarshal(body, request)
	if err != nil {
		status.BadRequest(w, err)
		return
	}

	services.Mine(request)
	if err != nil {
		status.ServerError(w, err)
		return
	}

	status.Success(w, []byte("OK"))
}
