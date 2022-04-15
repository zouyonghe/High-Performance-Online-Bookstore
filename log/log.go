package log

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
)

type Logger struct {
	Log *zap.Logger
}

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

	//core := zapcore.NewCore(encoder, writeSyncer, logLevel)
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, writeSyncer, logLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), logLevel),
	)

	logger := zap.New(core)
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			log.Fatal(err)
		}
	}(logger)
	return logger
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   viper.GetString("log.log_file"), //日志文件存放目录
		MaxSize:    viper.GetInt("log.max_size"),    //文件大小限制,单位MB
		MaxBackups: viper.GetInt("log.max_backups"), //最大保留日志文件数量
		MaxAge:     viper.GetInt("log.max_age"),     //日志文件保留天数
		Compress:   viper.GetBool("log.compress"),   //压缩                             //是否压缩处理
	})
}
