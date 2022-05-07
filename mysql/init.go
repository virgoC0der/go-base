package mysql

import (
	"context"
	"fmt"
	"time"

	. "github.com/virgoC0der/go-base/logging"
	"go.uber.org/zap"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlOption struct {
	Addr      string `ini:"addr"`       // 地址
	Port      int    `ini:"port"`       // 端口
	Username  string `ini:"username"`   // 用户名
	Password  string `ini:"password"`   // 密码
	DefaultDB string `ini:"default_db"` // 默认数据库名
	MaxIdle   int    `ini:"max_idle"`   // 最大空闲时间
	MaxConn   int    `ini:"max_conn"`   // 最大连接数
	Timeout   int    `ini:"timeout"`    // 超时时间
}

const (
	ConfPath         = "/home/go-base/conf/mysql.ini"
	MysqlUrlTemplate = "%s:%s@tcp(%s:%d)/%s"
	MysqlSuffix      = "?charset=utf8&parseTime=true&loc=Asia%2FShanghai"
)

var (
	globalDB *gorm.DB
	Timeout  time.Duration
)

func Init() error {
	conf, err := ini.Load(ConfPath)
	if err != nil {
		Logger.Warn("load mysql config failed, err", zap.Error(err))
		return err
	}

	opt := &MysqlOption{}
	if err = conf.MapTo(opt); err != nil {
		Logger.Warn("map mysql config to struct failed, err", zap.Error(err))
		return err
	}

	Timeout = time.Duration(opt.Timeout) * time.Second

	url := fmt.Sprintf(MysqlUrlTemplate, opt.Username, opt.Password, opt.Addr, opt.Port, opt.DefaultDB)
	globalDB, err = gorm.Open(mysql.Open(url+MysqlSuffix), &gorm.Config{})
	if err != nil {
		Logger.Warn("open mysql failed, err", zap.Error(err))
		return err
	}

	db, err := globalDB.DB()
	if err != nil {
		Logger.Warn("get db failed, err", zap.Error(err))
		return err
	}

	db.SetMaxIdleConns(opt.MaxIdle)
	db.SetMaxOpenConns(opt.MaxConn)
	return nil
}

// GetDBWithContext returns gorm.DB instance with context
func GetDBWithContext(ctx context.Context) *gorm.DB {
	return globalDB.WithContext(ctx)
}

// GetDB returns gorm.DB instance
func GetDB() *gorm.DB {
	return globalDB
}
