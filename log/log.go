package log

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

// InitLogger inits a logger
func InitLogger() *zap.Logger {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	var logLevel zapcore.Level

	switch viper.GetString("log.log_level") {
	case "DEBUG":
		logLevel = zapcore.DebugLevel
	case "INFO":
		logLevel = zapcore.InfoLevel
	case "WARN":
		logLevel = zapcore.WarnLevel
	case "ERROR":
		logLevel = zapcore.ErrorLevel
	case "DPANIC":
		logLevel = zapcore.DPanicLevel
	case "PANIC":
		logLevel = zapcore.PanicLevel
	case "FATAL":
		logLevel = zapcore.FatalLevel
	default:
		logLevel = zapcore.DebugLevel
	}
	// Default output log in stdout
	core := zapcore.NewCore(encoder, os.Stdout, logLevel)
	// Optional output log in file
	if viper.GetBool("log.file_output") {
		core = zapcore.NewTee(
			core,
			zapcore.NewCore(encoder, writeSyncer, logLevel),
		)
	}

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)
	return logger
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

// getLogWriter returns a writeSyncer
func getLogWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   viper.GetString("log.log_file"), //log file path
		MaxSize:    viper.GetInt("log.max_size"),    //file size limit
		MaxBackups: viper.GetInt("log.max_backups"), //max number of backups
		MaxAge:     viper.GetInt("log.max_age"),     //max age limit
		Compress:   viper.GetBool("log.compress"),   //log compression mode
	})
}
