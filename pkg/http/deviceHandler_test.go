package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pklimuk-eng-thesis/device/pkg/domain"
	"github.com/pklimuk-eng-thesis/device/pkg/service"
	"github.com/stretchr/testify/assert"
)

func TestDeviceHandler_GetInfo(t *testing.T) {
	deviceService := new(service.MockDeviceService)
	expectedDevice := &domain.Device{Enabled: true}
	deviceService.EXPECT().GetInfo().Return(expectedDevice)

	handler := NewDeviceHandler(deviceService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	handler.GetInfo(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"enabled": true}`, w.Body.String())
}

func TestDeviceHandler_ToggleEnabled(t *testing.T) {
	deviceService := new(service.MockDeviceService)
	expectedDevice := &domain.Device{Enabled: true}
	deviceService.EXPECT().ToggleEnabled().Return(expectedDevice)

	handler := NewDeviceHandler(deviceService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	handler.ToggleEnabled(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"enabled": true}`, w.Body.String())
}
