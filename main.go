package main

import (
	"High-Performance-Online-Bookstore/config"
	"High-Performance-Online-Bookstore/database"
	"High-Performance-Online-Bookstore/permission"
	ver "High-Performance-Online-Bookstore/pkg/version"
	"High-Performance-Online-Bookstore/router"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
)

var (
	cfg     = pflag.StringP("config", "c", "", "Specify config file path.")
	version = pflag.BoolP("version", "v", false, "show version info.")
	logger  *zap.Logger
)

// @title           High-Performance-Online-Bookstore
// @version         0.7.0
// @description     The jinshuzhai bookstore api server.
// @termsOfService  https://github.com/zouyonghe

// @contact.name   API Support
// @contact.url    https://github.com/zouyonghe
// @contact.email  1259085392z@gmail.com

// @license.name  GPLv3
// @license.url   https://www.gnu.org/licenses/gpl-3.0.html

// @host      127.0.0.1:8081
// @BasePath  /v1
// @Schemes   http https
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
	database.Init()

	// initialize RBAC
	permission.Init()

	// initialize router
	router.Init()

}
