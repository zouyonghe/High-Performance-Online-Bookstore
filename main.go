package main

import (
	"Jinshuzhai-Bookstore/config"
	"Jinshuzhai-Bookstore/model"
	ver "Jinshuzhai-Bookstore/pkg/version"
	"Jinshuzhai-Bookstore/router"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
)

var (
	cfg     = pflag.StringP("config", "c", "", "Specify config file path.")
	version = pflag.BoolP("version", "v", false, "show version info.")

	logger *zap.Logger
)

func main() {
	// 解析命令行参数
	pflag.Parse()

	// 显示版本信息
	ver.CheckShowVersion(*version)

	// 初始化配置
	logger = config.Init(*cfg)

	model.DB.Init(logger)

	router.InitRouter(logger)

	logger.Info("Jinshuzhai-Bookstore service started")
}
