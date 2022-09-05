package main

import (
	"context"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

type transactionKey struct{}

var trxKey = &transactionKey{}

func GetDb(ctx context.Context) *gorm.DB {
	if db, ok := ctx.Value(trxKey).(*gorm.DB); ok {
		return db
	}
	return db.WithContext(ctx)
}

func Transaction(ctx context.Context, fc func(ctx context.Context) error) error {
	return GetDb(ctx).Transaction(func(tx *gorm.DB) error {
		return fc(context.WithValue(ctx, trxKey, tx.WithContext(ctx)))
	})
}
