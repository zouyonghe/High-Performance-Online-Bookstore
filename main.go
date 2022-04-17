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
	// parse command line arguments
	pflag.Parse()

	// show version info
	ver.CheckShowVersion(*version)

	// initialize config
	logger = config.Init(*cfg)

	// initialize gorm
	model.DB.Init(logger)

	// initialize router
	router.InitRouter(logger)

	logger.Info("Jinshuzhai-Bookstore service started")
}
