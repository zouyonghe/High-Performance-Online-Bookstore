package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"strings"
)

type Config struct {
	Name string
}

// Init
func Init(cfg string, log *zap.Logger) error {
	c := Config{
		Name: cfg,
	}

	if err := c.initConfig(log); err != nil {
		return err
	}

	//Logger, _ := c.initLogger()

	if err := c.watchConfig(log); err != nil {
		return err
	}

	return nil
}

// initConfig 初始配置设置
func (c *Config) initConfig(log *zap.Logger) error {
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
	log.Info("Using config file: " + viper.ConfigFileUsed())
	return nil
}

// watchConfig 监听配置文件变化
func (c *Config) watchConfig(log *zap.Logger) error {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Info("Config file changed:", zap.String("change", e.Name))
	})
	return nil
}
