package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pklimuk-eng-thesis/device/pkg/service"
)

type DeviceHandler struct {
	DeviceService service.DeviceService
}

func NewDeviceHandler(deviceService service.DeviceService) *DeviceHandler {
	return &DeviceHandler{DeviceService: deviceService}
}

func (h *DeviceHandler) GetInfo(c *gin.Context) {
	device := h.DeviceService.GetInfo()
	c.IndentedJSON(http.StatusOK, &device)
}

func (h *DeviceHandler) ToggleEnabled(c *gin.Context) {
	device := h.DeviceService.ToggleEnabled()
	c.IndentedJSON(http.StatusOK, &device)
}
