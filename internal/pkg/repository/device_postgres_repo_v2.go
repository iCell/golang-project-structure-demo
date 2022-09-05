package repository

import (
	"context"
	"demo/internal/pkg/common"
	"demo/internal/pkg/domain"
)

type DeviceAddressTrxRepo struct {
}

type DevicePostgresRepoV2 struct {
	db common.DbContextGetter
}

func NewDevicePostgresRepoV2(db common.DbContextGetter) domain.IDeviceStore {
	return &DevicePostgresRepoV2{db: db}
}

func (d *DevicePostgresRepoV2) GetDevice(ctx context.Context, deviceID string) (domain.Device, error) {
	var device postgresDevice
	if err := d.db.GetDb(ctx).First(&device, "guid = ?", deviceID).Error; err != nil {
		return domain.Device{}, err
	}
	return device.Device, nil
}

func (d *DevicePostgresRepoV2) DeleteDevice(ctx context.Context, deviceID string) (domain.Device, error) {
	// suppose we need a trx here
	d.db.Transaction(ctx, func(ctx context.Context) error {
		_, err := d.GetDevice(ctx, deviceID)
		// ....
		// ....
		return err
	})
	return domain.Device{}, nil
}

func (d *DevicePostgresRepoV2) CreateDevice(ctx context.Context, device domain.Device) (domain.Device, error) {
	return domain.Device{}, nil
}

func (d *DevicePostgresRepoV2) DisableDeviceNotificationsForToken(ctx context.Context, token string) error {
	return nil
}

func (d *DevicePostgresRepoV2) GetInactiveDevices(ctx context.Context, inactiveDays, fromDeviceId uint, limit int) ([]domain.Device, error) {
	return nil, nil
}
