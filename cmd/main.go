package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pklimuk-eng-thesis/device/pkg/domain"
	dHttp "github.com/pklimuk-eng-thesis/device/pkg/http"
	dService "github.com/pklimuk-eng-thesis/device/pkg/service"
)

func main() {
	device := domain.Device{Enabled: true}

	deviceService := dService.NewDeviceService(&device)
	deviceHandler := dHttp.NewDeviceHandler(deviceService)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	dHttp.SetupRouter(r, deviceHandler)

	serviceAddress := os.Getenv("ADDRESS")
	if serviceAddress == "" {
		serviceAddress = ":8084"
	}
	log.Fatal(r.Run(serviceAddress))
}
