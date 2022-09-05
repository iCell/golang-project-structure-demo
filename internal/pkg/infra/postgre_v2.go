package infra

import (
	"context"
	"gorm.io/gorm"
)

type transactionKey struct{}

var trxKey = &transactionKey{}

type WrappedDb struct {
	globalDb *gorm.DB
}

// nested transaction
func (wd *WrappedDb) GetDb(ctx context.Context) *gorm.DB {
	if db, ok := ctx.Value(trxKey).(*gorm.DB); ok {
		return db
	}
	return wd.globalDb.WithContext(ctx)
}

func (wd *WrappedDb) Transaction(ctx context.Context, fc func(ctx context.Context) error) error {
	return wd.GetDb(ctx).Transaction(func(tx *gorm.DB) error {
		return fc(context.WithValue(ctx, trxKey, tx))
	})
}
