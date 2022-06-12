package g2gorm

import (
	"context"
	"database/sql"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/dijkvy/geocentric/g2gorm/config"
)

type ConfigOption func(c *gorm.Config)

func WithName(strategy schema.Namer) ConfigOption {
	return func(c *gorm.Config) {
		c.NamingStrategy = strategy
	}
}

func WithLogger(l logger.Interface) ConfigOption {
	return func(c *gorm.Config) {
		c.Logger = l
	}
}

func WithSkipDefaultTransaction(skip bool) ConfigOption {
	return func(c *gorm.Config) {
		c.SkipDefaultTransaction = skip
	}
}

func WithPrepareStmt(preStmt bool) ConfigOption {
	return func(c *gorm.Config) {
		c.PrepareStmt = preStmt
	}
}

func WithDisableAutomaticPing(autoPing bool) ConfigOption {
	return func(c *gorm.Config) {
		c.DisableAutomaticPing = autoPing
	}
}

func WithCreateBatchSize(sz int) ConfigOption {
	return func(c *gorm.Config) {
		c.CreateBatchSize = sz
	}
}

func WithSingularTable() ConfigOption {
	return WithName(schema.NamingStrategy{SingularTable: true})
}

type dbEntry struct {
	db *gorm.DB
}

func (c *dbEntry) WithContext(ctx context.Context) (db *gorm.DB) {
	return c.db.WithContext(ctx)
}

func NewDBHelper(conf *config.EngineConfig, opts ...ConfigOption) (helper DBHelper, err error) {
	var dialect DialectFunc
	if dialect, err = getDialect(conf.GetDriver()); err != nil {
		return nil, err
	}

	var db *gorm.DB
	var dbConfig gorm.Config
	for _, fn := range opts {
		fn(&dbConfig)
	}

	if db, err = gorm.Open(dialect(conf.GetDsn()), &dbConfig); err != nil {
		return nil, err
	}

	{
		var sqlDB *sql.DB
		if sqlDB, err = db.DB(); err != nil {
			return nil, err
		}
		if conf.MaxOpenConn != nil {
			sqlDB.SetMaxOpenConns(int(conf.GetMaxOpenConn()))
		}

		if conf.MaxIdleConn != nil {
			sqlDB.SetMaxIdleConns(int(conf.GetMaxIdleConn()))
		}

		if conf.MaxLifeTime != nil {
			sqlDB.SetConnMaxLifetime(time.Duration(conf.GetMaxLifeTime()) * time.Second)
		}

		if conf.MaxIdleTime != nil {
			sqlDB.SetConnMaxIdleTime(time.Duration(conf.GetMaxIdleTime()) * time.Second)
		}

	}
	return &dbEntry{db: db}, err
}
