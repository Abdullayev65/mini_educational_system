package app

import (
	"context"
	"database/sql"
	"github.com/Abdullayev65/mini_educational_system/pkg/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

var resetModel = false

func (a *App) database() (*bun.DB, context.Context) {
	dsn := "postgres:root123//postgres:localhost:5432/postgres?sslmode=disable"
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(pgdb, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	ctx := context.Background()

	ms := []interface{}{
		(*models.Department)(nil),
		(*models.Level)(nil),
		(*models.Subject)(nil),
		(*models.User)(nil),
		(*models.Attendance)(nil),
		(*models.Position)(nil),
	}
	if resetModel {
		db.ResetModel(ctx, ms...)
	} else {
		db.RegisterModel(ms...)
	}
	return db, ctx

}
