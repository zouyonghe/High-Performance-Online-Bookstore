package model

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Database struct {
	Self   *gorm.DB
	Docker *gorm.DB
}

var DB *Database

// openDB is used to open a database connection.
func openDB(username, password, addr, name string, logger *zap.Logger) *gorm.DB {
	DSN := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
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
	cfg := &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tb_", //表名前缀
			SingularTable: true,  //是否复数表名
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	db, err := gorm.Open(sql, cfg)
	if err != nil {
		logger.Error("Database connection failed. Database name: %s", zap.Error(err))
	}

	return db
}

// InitSelfDB initialize the database connection.
func InitSelfDB(logger *zap.Logger) *gorm.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"),
		logger)
}

// GetSelfDB return the global database connection.
func GetSelfDB(logger *zap.Logger) *gorm.DB {
	return InitSelfDB(logger)
}

// InitDockerDB initialize the docker database connection.
func InitDockerDB(logger *zap.Logger) *gorm.DB {
	return openDB(viper.GetString("docker_db.username"),
		viper.GetString("docker_db.password"),
		viper.GetString("docker_db.addr"),
		viper.GetString("docker_db.name"),
		logger)
}

// GetDockerDB return the docker database connection.
func GetDockerDB(logger *zap.Logger) *gorm.DB {
	return InitDockerDB(logger)
}

// Init initialize the database connection.
func (db *Database) Init(logger *zap.Logger) {
	DB = &Database{
		Self:   GetSelfDB(logger),
		Docker: GetDockerDB(logger),
	}
	logger.Info("Database connection established.")
}

/*func (db *Database) Close() {
	DB.Self.Close()
	DB.Docker.Close()
}*/
