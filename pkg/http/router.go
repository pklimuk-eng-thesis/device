package http

import "github.com/gin-gonic/gin"

var getInfoEndpoint = "/info"
var enabledEndpoint = "/enabled"

func SetupRouter(r *gin.Engine, h *DeviceHandler) {
	r.GET(getInfoEndpoint, h.GetInfo)
	r.PATCH(enabledEndpoint, h.ToggleEnabled)
}
