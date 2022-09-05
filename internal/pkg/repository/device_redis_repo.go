package repository

import (
	"context"
	"demo/internal/pkg/domain"
	"github.com/trustwallet/go-libs/cache/redis"
	"time"
)

type DeviceRedisRepo struct {
	redis *redis.Redis
}

func NewDeviceRedisRepo(redis *redis.Redis) domain.IDeviceManager {
	return &DeviceRedisRepo{redis: redis}
}

func (d *DeviceRedisRepo) GetDevice(ctx context.Context, deviceID string) (domain.Device, error) {
	return domain.Device{}, nil
}

func (d *DeviceRedisRepo) CreateDevice(ctx context.Context, device domain.Device) (domain.Device, error) {
	if err := d.redis.Set(ctx, device.GUID, &device, time.Minute); err != nil {
		return domain.Device{}, err
	}
	return device, nil
}

func (d *DeviceRedisRepo) DeleteDevice(ctx context.Context, deviceID string) (domain.Device, error) {
	return domain.Device{}, nil
}
