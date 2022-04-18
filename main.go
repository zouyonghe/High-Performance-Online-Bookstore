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
	logger  *zap.Logger
)

func main() {

	// parse command line arguments
	pflag.Parse()

	// show version info
	ver.CheckShowVersion(*version)

	// initialize config
	logger = config.Init(*cfg)

	// Replace global logger
	gl := zap.ReplaceGlobals(logger)
	defer gl()

	// initialize gorm
	model.DB.Init()

	// initialize router
	router.InitRouter()

}
