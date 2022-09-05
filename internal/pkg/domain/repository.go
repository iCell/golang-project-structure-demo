package domain

import "context"

type IDeviceStore interface {
	IDeviceManager
	DisableDeviceNotificationsForToken(ctx context.Context, token string) error
	GetInactiveDevices(ctx context.Context, inactiveDays, fromDeviceId uint, limit int) ([]Device, error)
}

type IDeviceManager interface {
	GetDevice(ctx context.Context, deviceID string) (Device, error)
	CreateDevice(ctx context.Context, device Device) (Device, error)
	DeleteDevice(ctx context.Context, deviceID string) (Device, error)
}
