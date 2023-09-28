package controllers

import (
	"net/http"

	"github.com/victormagalhaess/go-api-template/pkg/api/status"
)

// @Description Healthcheck endpoint
// @Summary Endpoint to check the health of the application
// @Produce  text/plain
// @Router /api/v1/healthcheck [GET]
// @Success 200
// @Tags Healthcheck
func Healthcheck(w http.ResponseWriter, r *http.Request) {
	status.Success(w, []byte("OK"))
}