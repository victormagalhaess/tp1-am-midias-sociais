package controllers_test

import (
	"net/http"
	"testing"

	"github.com/victormagalhaess/go-api-template/pkg/api/controllers"
	"github.com/victormagalhaess/go-api-template/pkg/mock"
)

func TestHealthcheck(t *testing.T) {
	w := &mock.ResponseWriter{}
	controllers.Healthcheck(w, nil)
	if w.Status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Status)
	}
	if string(w.Data) != "OK" {
		t.Errorf("Expected data %s, got %s", "OK", w.Data)
	}
}