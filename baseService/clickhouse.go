package baseService

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/juntao3000/fxh-spider-go-common/baseCommon"
	gormCh "gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	ClickHouseClient driver.Conn
	ClickHouseGormDb *gorm.DB
)

func DisposeClickHouse() {
	if ClickHouseClient != nil {
		_ = ClickHouseClient.Close()
	}
}

func getClickHouseOptions() *clickhouse.Options {
	opt := &clickhouse.Options{
		Protocol: clickhouse.Protocol(baseCommon.BaseConfig.ClickHouseProtocol),
		Addr:     []string{fmt.Sprintf("%s:%d", baseCommon.BaseConfig.ClickHouseHost, baseCommon.BaseConfig.ClickHousePort)},
		Auth: clickhouse.Auth{
			Username: baseCommon.BaseConfig.ClickHouseUser,
			Password: baseCommon.BaseConfig.ClickHousePassword,
		},
		//Compression: &clickhouse.Compression{
		//	Method: clickhouse.CompressionZSTD,
		//},
		ConnOpenStrategy: clickhouse.ConnOpenRoundRobin,
		Settings: clickhouse.Settings{
			"allow_experimental_lightweight_delete": 1,
			"allow_experimental_object_type":        1,
		},
	}
	if len(baseCommon.BaseConfig.ClickHouseDatabase) > 0 {
		opt.Auth.Database = baseCommon.BaseConfig.ClickHouseDatabase
	}
	//if baseCommon.BaseConfig.ClickHouseMaxOpenConnection > 1 {
	//	opt.MaxOpenConns = baseCommon.BaseConfig.ClickHouseMaxOpenConnection
	//}

	return opt
}

func InitClickHouse() error {
	if ClickHouseClient != nil {
		return nil
	}

	opt := getClickHouseOptions()
	conn, err := clickhouse.Open(opt)
	if err != nil {
		return err
	}

	err = conn.Ping(context.Background())
	if err != nil {
		_ = conn.Close()
		return err
	}

	gormDb, err := GetClickHouseGormDb()
	if err != nil {
		_ = conn.Close()
		return err
	}

	ClickHouseClient = conn
	ClickHouseGormDb = gormDb
	return nil
}

func GetClickHouseGormDb() (*gorm.DB, error) {
	if ClickHouseGormDb != nil {
		return ClickHouseGormDb, nil
	}

	opt := getClickHouseOptions()
	sqlDb := clickhouse.OpenDB(opt)
	dialector := gormCh.New(gormCh.Config{Conn: sqlDb})

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
	gormCfg := &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 newLogger,
	}

	return gorm.Open(dialector, gormCfg)
}
