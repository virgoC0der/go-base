package rbac

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"go.uber.org/zap"
	"gopkg.in/ini.v1"
	"time"

	. "github.com/virgoC0der/go-base/logging"
	"github.com/virgoC0der/go-base/mysql"
)

var (
	adapter  *gormadapter.Adapter
	enforcer *casbin.SyncedEnforcer
)

// NewAdapter returns a GormAdapter instance
func NewAdapter() error {
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

	enforcer, err = casbin.NewSyncedEnforcer("/Users/chensx/Desktop/Go/go-base/rbac/conf/model.conf", adapter)
	if err != nil {
		Logger.Warn("new enforcer failed", zap.Error(err))
		return err
	}
	enforcer.EnableEnforce(true)
	_ = enforcer.LoadPolicy()
	enforcer.StartAutoLoadPolicy(10 * time.Second)

	return nil
}

type chCasbinPolicyItem struct {
	ctx context.Context
	e   *casbin.SyncedEnforcer
}

var ch chan *chCasbinPolicyItem

func init() {
	ch = make(chan *chCasbinPolicyItem, 1)
	// hot reload casbin policy
	go func() {
		for item := range ch {
			if err := item.e.LoadPolicy(); err != nil {
				Logger.Warn("load policy failed", zap.Error(err))
			}
		}
	}()
}

// LoadCasbinPolicy async load casbin policy
func LoadCasbinPolicy(ctx context.Context, e *casbin.SyncedEnforcer) {
	if len(ch) > 0 {
		Logger.Info("casbin policy is loading, skip")
		return
	}

	ch <- &chCasbinPolicyItem{ctx: ctx, e: e}
}

// RegisterFunc registers a function to the enforcer
func RegisterFunc(name string, f func(args ...interface{}) (interface{}, error)) {
	enforcer.AddFunction(name, f)
}

func Create(rule *mysql.CasbinRule) (bool, error) {
	return enforcer.AddPolicy(rule.RoleId, rule.Api, rule.Method)
}

func List(roleId string) [][]string {
	return enforcer.GetFilteredPolicy(0, roleId)
}

func Enforce(sub, dom, obj, act string) (bool, error) {
	return enforcer.Enforce(sub, dom, obj, act)
}
