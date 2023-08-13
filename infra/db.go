package infra

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"time"

	"canaanadvisors-test/config"
)

const (
	psqlConn = "host=%v user=%v dbname=%v sslmode=disable password=%v port=%v search_path=%v"
)

func NewDB(ctx context.Context, logger *zap.Logger) (*gorm.DB, error) {
	strConn := fmt.Sprintf(psqlConn, config.C.Database.Host, config.C.Database.User,
		config.C.Database.DBName, config.C.Database.Password, config.C.Database.Port, config.C.Database.Schema)
	db, err := gorm.Open(postgres.Open(strConn), &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Silent),
	})
	if err == nil {
		if sqlDB, e := db.DB(); e == nil {
			sqlDB.SetMaxIdleConns(config.C.Database.MaxIdleConn)
			sqlDB.SetMaxOpenConns(config.C.Database.MaxOpenConn)
			sqlDB.SetConnMaxLifetime(time.Duration(config.C.Database.ConnMaxLifetimeSecond) * time.Second)
		}
	} else {
		logger.Error(err.Error())
	}
	return db, err
}
