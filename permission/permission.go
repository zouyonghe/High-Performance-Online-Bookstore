package permission

import (
	"High-Performance-Online-Bookstore/log"
	"github.com/casbin/casbin/v2"
	zaplogger "github.com/casbin/zap-logger/v2"
	"go.uber.org/zap"
)

var E *casbin.Enforcer

// InitPermission inits the permission.
func InitPermission() {
	logger := zaplogger.NewLoggerByZap(zap.L(), true)
	var err error
	E, err = casbin.NewEnforcer("conf/model.conf", "conf/policy.csv")
	if err != nil {
		zap.L().Error("init permission error", zap.Error(err))
		return
	}
	if err = E.LoadPolicy(); err != nil {
		zap.L().Error("load permission error", zap.Error(err))
		return
	}
	E.EnableLog(true)
	E.SetLogger(logger)
}

func CheckPermission(sub, obj, act string) bool {
	ok, err := E.Enforce(sub, obj, act)
	if err != nil {
		log.ErrCheckPermission(err)
		return false
	}
	return ok
}
