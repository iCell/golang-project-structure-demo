package repository

import (
	"context"
	"demo/internal/pkg/domain"
	"gorm.io/gorm"
)

type postgresDevice struct {
	gorm.Model
	domain.Device
}

type DevicePostgresRepo struct {
	db *gorm.DB
}

/*
1. which level to put the transaction code.

1.1 manage trx in the use_case: 1)depend on the gorm; 2) define a new interface (win)

1.2 use trx in the repository level:

wallet + device + address within one single trx: 3 repo within 1 trx
*/

func NewDevicePostgresRepo(db *gorm.DB) domain.IDeviceStore {
	return &DevicePostgresRepo{db: db}
}

func (d *DevicePostgresRepo) GetDevice(ctx context.Context, deviceID string) (domain.Device, error) {
	return d.getDevice(d.db.WithContext(ctx), deviceID)
}

func (d *DevicePostgresRepo) getDevice(tx *gorm.DB, deviceID string) (domain.Device, error) {
	var device postgresDevice
	if err := tx.First(&device, "guid = ?", deviceID).Error; err != nil {
		return domain.Device{}, err
	}
	return device.Device, nil
}

func (d *DevicePostgresRepo) DeleteDevice(ctx context.Context, deviceID string) (domain.Device, error) {
	// suppose we need a trx here
	d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		_, err := d.getDevice(tx, deviceID)
		// ....
		// ....
		return err
	})
	return domain.Device{}, nil
}

func (d *DevicePostgresRepo) CreateDevice(ctx context.Context, device domain.Device) (domain.Device, error) {
	return domain.Device{}, nil
}

func (d *DevicePostgresRepo) DisableDeviceNotificationsForToken(ctx context.Context, token string) error {
	return nil
}

func (d *DevicePostgresRepo) GetInactiveDevices(ctx context.Context, inactiveDays, fromDeviceId uint, limit int) ([]domain.Device, error) {
	return nil, nil
}
