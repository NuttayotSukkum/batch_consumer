package db_config

import (
	"context"
	"fmt"
	"github.com/NuttayotSukkum/batch_consumer/internals/models/dao"
	logger "github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
)

func ACNDBConnector(ctx context.Context, host, port, username, password, dbName string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)
	logger.Infof("%v: database connect:%+v", ctx, dsn)
	dial := mysql.Open(dsn)
	db, err := gorm.Open(dial, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: gormlogger.Discard,
	})
	if err != nil {
		log.Panic(ctx, err.Error())
	}
	if err := db.AutoMigrate(&dao.BatchHeader{}); err != nil {
		logger.Errorf("unable to migrate model: %s", err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		log.Panicf("%v :failed to create connection pools :%s", ctx, err)
	}
	logger.Infof("Database is running!... %v db_stats %s", ctx, sqlDb.Stats())
	return db
}
