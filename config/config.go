package config

import (
	"High-Performance-Online-Bookstore/log"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	l "log"
	"strings"
)

type Config struct {
	Name string
}

// Init inits configures and deal with errors
// then return a zap logger.
func Init(cfg string) *zap.Logger {
	logger, err := initAll(cfg)
	if err != nil {
		l.Fatal("init config failed", err)
	}
	return logger
}

// initAll inits config and log,
// and it will watch config file change.
func initAll(cfg string) (*zap.Logger, error) {
	c := Config{
		Name: cfg,
	}
	// initialize log file
	if err := c.initConfig(); err != nil {
		return nil, err
	}

	// build a logger example
	logger := log.InitLogger()

	// watch config file change
	if err := c.watchConfig(); err != nil {
		return nil, err
	}

	return logger, nil
}

// initConfig read the config file
// and initialize the config.
func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("SERVER")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// watchConfig watches the config file change.
func (c *Config) watchConfig() error {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		zap.L().Info("Config file changed:" + e.Name)
	})
	return nil
}
