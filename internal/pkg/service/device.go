package service

import (
	"context"
	"demo/internal/pkg/domain"
	"errors"
	"github.com/trustwallet/go-libs/cache/redis"
)

type DeviceUseCase struct {
	redisRepo    domain.IDeviceManager
	postgresRepo domain.IDeviceStore
}

func NewDeviceUseCase(redisRepo domain.IDeviceManager, postgresRepo domain.IDeviceStore) *DeviceUseCase {
	return &DeviceUseCase{
		redisRepo:    redisRepo,
		postgresRepo: postgresRepo,
	}
}

func (du *DeviceUseCase) GetDevice(ctx context.Context, deviceId string) (domain.Device, error) {
	device, err := du.redisRepo.GetDevice(ctx, deviceId)
	if errors.Is(err, redis.ErrNotFound) {
		device, err = du.postgresRepo.GetDevice(ctx, deviceId)
	}
	if err != nil {
		return domain.Device{}, err
	}
	return device, nil
}
