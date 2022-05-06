package log

import "go.uber.org/zap"

// router

func StartListenHTTP(addr string) {
	zap.L().Info("start listening HTTP request", zap.String("addr", addr))
}

func StartListenHTTPS(addr string) {
	zap.L().Info("start listening HTTPS requests", zap.String("addr", addr))
}

func WaitForRouter() {
	zap.L().Info("Waiting for the router, retry in 1 second.")
}

func RouterDeployed() {
	zap.L().Info("The router has been deployed successfully.")
}
