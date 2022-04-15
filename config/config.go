package config

import (
	"Jinshuzhai-Bookstore/log"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	l "log"
	"strings"
)

type Config struct {
	Name string
}

// Init 初始化配置
func Init(cfg string) *zap.Logger {
	logger, err := initAll(cfg)
	if err != nil {
		l.Fatal("init config failed", zap.Error(err))
	}
	return logger
}

// initAll 初始化配置和日志
func initAll(cfg string) (*zap.Logger, error) {
	c := Config{
		Name: cfg,
	}
	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return nil, err
	}

	// 建立logger实例
	logger := log.InitLogger()

	// 监控配置文件变化并热加载程序
	if err := c.watchConfig(logger); err != nil {
		return nil, err
	}

	return logger, nil
}

// initConfig 初始配置设置
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

// watchConfig 监听配置文件变化
func (c *Config) watchConfig(logger *zap.Logger) error {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.Info("Config file changed:" + e.Name)
	})
	return nil
}
