package common

import (
	"context"
	"gorm.io/gorm"
)

type DbContextGetter interface {
	GetDb(ctx context.Context) *gorm.DB
	Transaction(ctx context.Context, fc func(ctx context.Context) error) error
}
