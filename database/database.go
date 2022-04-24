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

	gormLogger := zapgorm2.New(zap.L())
	gormLogger.SetAsDefault() // optional: configure gorm to use this zapgorm.Logger for callbacks

	cfg := &gorm.Config{
		SkipDefaultTransaction: false,
		Logger:                 gormLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tb_", //表名前缀
			SingularTable: true,  //是否复数表名
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	db, err := gorm.Open(sql, cfg)
	if err != nil {
		zap.L().Error("Database connection failed. Database name: %s", zap.Error(err))
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
func (db *Database) InitDatabase() {
	DB = &Database{
		Self:   GetSelfDB(),
		Docker: GetDockerDB(),
	}
	zap.L().Info("Database connection established.")
}
