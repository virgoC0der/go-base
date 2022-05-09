package rbac

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
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
func NewAdapter() error {
	conf, err := ini.Load("./conf/mysql.ini")
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

	enforcer, err = casbin.NewEnforcer("./conf/model.conf", adapter)
	if err != nil {
		Logger.Warn("new enforcer failed", zap.Error(err))
		return err
	}
	_ = enforcer.LoadPolicy()

	return nil
}

func Create(rule *mysql.CasbinRule) (bool, error) {
	return enforcer.AddPolicy(rule.RoleId, rule.Api, rule.Method)
}

func List(roleId string) [][]string {
	return enforcer.GetFilteredPolicy(0, roleId)
}

func Enforcer(sub, obj, act string) (bool, error) {
	return enforcer.Enforce(sub, obj, act)
}
