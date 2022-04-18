package test

import "go.uber.org/zap"

func Test() {
	zap.L().Info("In test")
}
