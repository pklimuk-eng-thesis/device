package service

import (
	"testing"

	"github.com/pklimuk-eng-thesis/device/pkg/domain"
	"github.com/stretchr/testify/assert"
)

func Test_DeviceService_NewDeviceService(t *testing.T) {
	device := &domain.Device{
		Enabled: true,
	}
	deviceService := NewDeviceService(device)
	assert.NotNil(t, deviceService)
}

func Test_DeviceService_IsEnabled(t *testing.T) {
	tests := []struct {
		name string
		s    deviceService
		want *domain.Device
	}{
		{
			name: "device is enabled",
			s:    deviceService{device: &domain.Device{Enabled: true}},
			want: &domain.Device{Enabled: true},
		},
		{
			name: "device is disabled",
			s:    deviceService{device: &domain.Device{Enabled: false}},
			want: &domain.Device{Enabled: false},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.s.GetInfo()
			assert.Equal(t, test.want.Enabled, got.Enabled)
		})
	}
}

func Test_DeviceService_ToggleEnabled(t *testing.T) {
	tests := []struct {
		name string
		s    deviceService
		want *domain.Device
	}{
		{
			name: "device is enabled",
			s:    deviceService{device: &domain.Device{Enabled: true}},
			want: &domain.Device{Enabled: false},
		},
		{
			name: "device is disabled",
			s:    deviceService{device: &domain.Device{Enabled: false}},
			want: &domain.Device{Enabled: true},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.s.ToggleEnabled()
			assert.Equal(t, test.want.Enabled, got.Enabled)
		})
	}
}
