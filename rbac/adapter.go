package rbac

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/jinzhu/gorm"
	. "github.com/virgoC0der/go-base/logging"
	"github.com/virgoC0der/go-base/mysql"
	"go.uber.org/zap"
)

var (
	adapter  *gormadapter.Adapter
	enforcer *casbin.Enforcer
)

// NewAdapter returns a GormAdapter instance
func NewAdapter(db *gorm.DB) error {
	var err error
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
