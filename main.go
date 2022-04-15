package main

import (
	"Jinshuzhai-Bookstore/config"
	ver "Jinshuzhai-Bookstore/pkg/version"
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
	logger, err := config.Init(*cfg)
	if err != nil {
		logger.Fatal("init config failed", zap.Error(err))
	}

	logger.Info("Start server successfully")
}
