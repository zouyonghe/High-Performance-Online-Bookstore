package policy

import (
	"github.com/casbin/casbin/v2"
	zaplogger "github.com/casbin/zap-logger/v2"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	E *casbin.Enforcer
)

// InitPolicy inits the policy.
func InitPolicy() {
	logger := zaplogger.NewLoggerByZap(zap.L(), true)
	var err error
	E, err = casbin.NewEnforcer("conf/model.conf", "conf/policy.csv")
	if err != nil {
		zap.L().Error("init policy error", zap.Error(err))
		return
	}
	if err := E.LoadPolicy(); err != nil {
		zap.L().Error("load policy error", zap.Error(err))
		return
	}
	E.EnableLog(true)
	E.SetLogger(logger)
}

func CheckPermission(ctx *gin.Context, sub, obj, act string) bool {
	zap.L().Info("checkPermission", zap.String("sub", sub), zap.String("obj", obj), zap.String("act", act))
	ok, err := E.Enforce(sub, obj, act)
	if err != nil {
		zap.L().Error("checkPermission error", zap.Error(err))
		return false
	}
	if !ok {
		zap.L().Error("checkPermission error", zap.String("sub", sub), zap.String("obj", obj), zap.String("act", act))
		return false
	}
	zap.L().Info("checkPermission ok", zap.String("sub", sub), zap.String("obj", obj), zap.String("act", act))
	return true
}