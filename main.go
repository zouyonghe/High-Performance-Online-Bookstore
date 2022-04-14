package main

import (
	"Jinshuzhai-Bookstore/config"
	"Jinshuzhai-Bookstore/log"
	"encoding/json"
	"fmt"
	//ver "github.com/hashicorp/go-version"
	ver "Jinshuzhai-Bookstore/pkg/version"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
	"os"
)

var (
	cfg     = pflag.StringP("config", "c", "", "Specify config file path.")
	version = pflag.BoolP("version", "v", false, "show version info.")
)
var logger *zap.Logger

func main() {
	logger = log.InitLogger()
	pflag.Parse()

	if *version {
		v := ver.Get()
		marshalled, err := json.MarshalIndent(&v, "", "  ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(marshalled))
		return
	}

	if err := config.Init(*cfg, logger); err != nil {
		panic(err)
	}
	//logger.Info("log", zap.String("log", "log"))
}
