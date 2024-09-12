package database

import (
	"database/sql"
	"os"
	"runtime"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

var Database *bun.DB

func Connect() error {
	config := pgdriver.NewConnector(
		pgdriver.WithAddr(os.Getenv("DATABASE_HOST")),
		pgdriver.WithUser(os.Getenv("DATABASE_USER")),
		pgdriver.WithPassword(os.Getenv("DATABASE_PASS")),
		pgdriver.WithDatabase(os.Getenv("DATABASE_NAME")),
		pgdriver.WithInsecure(true),
	)

	sqldb := sql.OpenDB(config)

	if err := sqldb.Ping(); err != nil {
		return err
	}

	Database = bun.NewDB(sqldb, pgdialect.New(), bun.WithDiscardUnknownColumns())

	maxOpenConns := 4 * runtime.GOMAXPROCS(0)

	Database.SetMaxOpenConns(maxOpenConns)
	Database.SetMaxIdleConns(maxOpenConns)
	Database.AddQueryHook(bundebug.NewQueryHook())

	return nil
}
