package database

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"moul.io/zapgorm2"
)

type Database struct {
	Self   *gorm.DB
	Docker *gorm.DB
}

var DB *Database

// openDB is used to open a database connection.
func openDB(username, password, addr, name string) *gorm.DB {
	DSN := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local")
	sql := mysql.New(mysql.Config{
		DSN:               DSN,
		DefaultStringSize: 256,
	})
	gormLogger := zapgorm2.New(zap.L())
	gormLogger.SetAsDefault()
	cfg := &gorm.Config{
		SkipDefaultTransaction: false,
		Logger:                 gormLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tb_",
			SingularTable: false,
		},
	}
	db, err := gorm.Open(sql, cfg)
	if err != nil {
		zap.L().Error("Database connection failed.", zap.Error(err))
		return nil
	}
	return db
}

// InitSelfDB initialize the database connection.
func InitSelfDB() *gorm.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"),
	)
}

// GetSelfDB return the global database connection.
func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

// InitDockerDB initialize the docker database connection.
func InitDockerDB() *gorm.DB {
	return openDB(viper.GetString("docker_db.username"),
		viper.GetString("docker_db.password"),
		viper.GetString("docker_db.addr"),
		viper.GetString("docker_db.name"),
	)
}

// GetDockerDB return the docker database connection.
func GetDockerDB() *gorm.DB {
	return InitDockerDB()
}

// InitDatabase initialize the database connection.
// The main database is required; the server exits if it
// cannot be connected. The docker database is optional
// and only connected when configured.
func (db *Database) InitDatabase() {
	self := GetSelfDB()
	if self == nil {
		zap.L().Fatal("Main database connection is required, exiting.")
	}

	DB = &Database{Self: self}
	if viper.GetString("docker_db.addr") != "" {
		DB.Docker = GetDockerDB()
		if DB.Docker == nil {
			zap.L().Warn("Docker database connection failed, continuing without it.")
		}
	}
	zap.L().Info("Database connection established.")
}

func Init() {
	DB.InitDatabase()
}
