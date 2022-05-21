package permission

import (
	"High-Performance-Online-Bookstore/log"
	"github.com/casbin/casbin/v2"
	zaplogger "github.com/casbin/zap-logger/v2"
	"go.uber.org/zap"
)

var E *casbin.Enforcer

func Init() {
	E = initPermission("conf/model.conf", "conf/policy.csv", true)
}

// InitPermission inits the permission.
func initPermission(modelPath string, policyPath string, log bool) (E *casbin.Enforcer) {
	E, err := casbin.NewEnforcer(modelPath, policyPath)
	if err != nil {
		zap.L().Error("init permission error", zap.Error(err))
		return
	}
	if err = E.LoadPolicy(); err != nil {
		zap.L().Error("load permission error", zap.Error(err))
		return
	}
	if log == true {
		logger := zaplogger.NewLoggerByZap(zap.L(), true)
		E.EnableLog(true)
		E.SetLogger(logger)
	}
	return E
}

func CheckPermission(sub, obj, act string) bool {
	ok, err := E.Enforce(sub, obj, act)
	if err != nil {
		log.ErrCheckPermission(err)
		return false
	}
	return ok
}
