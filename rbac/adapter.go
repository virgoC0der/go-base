package rbac

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"gopkg.in/ini.v1"

	. "github.com/virgoC0der/go-base/logging"
	"github.com/virgoC0der/go-base/mysql"
)

var (
	adapter  *gormadapter.Adapter
	enforcer *casbin.Enforcer
)

// NewAdapter returns a GormAdapter instance
func NewAdapter(db *gorm.DB) error {
	conf, err := ini.Load(mysql.ConfPath)
	if err != nil {
		Logger.Warn("load mysql config failed", zap.Error(err))
		return err
	}

	opt := &mysql.MysqlOption{}
	if err = conf.MapTo(opt); err != nil {
		Logger.Warn("map mysql config to struct failed", zap.Error(err))
		return err
	}

	adapter, err = gormadapter.NewAdapterByDB(db)
	if err != nil {
		Logger.Warn("new adapter failed", zap.Error(err))
		return err
	}

	enforcer, err = casbin.NewEnforcer("./conf/model.conf", adapter)
	if err != nil {
		Logger.Warn("new enforcer failed", zap.Error(err))
		return err
	}

	return nil
}
