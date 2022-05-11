package rbac

import (
	"fmt"
	"time"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/virgoC0der/go-base/mysql"
	"go.uber.org/zap"
	"gopkg.in/ini.v1"

	. "github.com/virgoC0der/go-base/logging"
)

func Init() error {
	conf, err := ini.Load("/Users/chensx/Desktop/Go/go-base/rbac/conf/mysql.ini")
	if err != nil {
		Logger.Warn("load mysql config failed", zap.Error(err))
		return err
	}

	opt := &mysql.MysqlOption{}
	if err = conf.MapTo(opt); err != nil {
		Logger.Warn("map mysql config to struct failed", zap.Error(err))
		return err
	}

	url := fmt.Sprintf(mysql.MysqlUrlTemplate, opt.Username, opt.Password, opt.Addr, opt.Port, opt.DefaultDB) + mysql.MysqlSuffix
	Logger.Info("mysql url", zap.String("url", url))
	adapter, err = gormadapter.NewAdapter("mysql", url, true)
	if err != nil {
		Logger.Warn("new adapter failed", zap.Error(err))
		return err
	}

	enforcer, err = casbin.NewSyncedEnforcer("/Users/chensx/Desktop/Go/go-base/rbac/conf/rbac_with_domains_model.conf", adapter)
	if err != nil {
		Logger.Warn("new enforcer failed", zap.Error(err))
		return err
	}
	enforcer.EnableEnforce(true)
	_ = enforcer.LoadPolicy()
	enforcer.StartAutoLoadPolicy(10 * time.Second)

	return nil
}
