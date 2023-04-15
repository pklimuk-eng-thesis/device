package service

import "github.com/pklimuk-eng-thesis/device/pkg/domain"

//go:generate --name DeviceService --output mock_deviceService.go
type DeviceService interface {
	GetInfo() *domain.Device
	ToggleEnabled() *domain.Device
}

type deviceService struct {
	device *domain.Device
}

func NewDeviceService(device *domain.Device) DeviceService {
	return &deviceService{device: device}
}

func (s *deviceService) GetInfo() *domain.Device {
	return s.device
}

func (s *deviceService) ToggleEnabled() *domain.Device {
	s.device.Enabled = !s.device.Enabled
	return s.device
}
