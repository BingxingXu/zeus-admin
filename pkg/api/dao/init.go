package dao

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/viper"
	"zeus/pkg/api/domain/search/adapter/statement"
	"zeus/pkg/api/log"
)

var (
	db *gorm.DB
)

const (
	DRIVER_PG = "pg"
	DRIVER_MYSQL  = "mysql"
	DRIVER_SQLITE = "sqlite"
)

var searchAdapter = &statement.SqlSearchAdapter{}

// Setup : Connect to mysql database
func Setup() {
	var err error
	switch viper.Get("database.driver") {
	case DRIVER_SQLITE:
		db, err = gorm.Open("sqlite3", viper.GetString("database.sqlite.dsn"))
		if err != nil {
			log.Fatal(fmt.Sprintf("Failed to connect sqlite %s", err.Error()))
		} else {
			db.LogMode(true)
		}
	case DRIVER_MYSQL:
		host := viper.GetString("database.mysql.host")
		user := viper.GetString("database.mysql.user")
		password := viper.GetString("database.mysql.password")
		name := viper.GetString("database.mysql.name")
		charset := viper.GetString("database.mysql.charset")

		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", user, password, host, name, charset)
		log.Debug(dsn)
		db, err = gorm.Open("mysql", dsn)
		if err != nil {
			log.Fatal(fmt.Sprintf("Failed to connect mysql %s", err.Error()))
		} else {
			db.DB().SetMaxIdleConns(viper.GetInt("database.mysql.pool.min"))
			db.DB().SetMaxOpenConns(viper.GetInt("database.mysql.pool.max"))
			if gin.Mode() != gin.ReleaseMode {
				db.LogMode(true)
			}
		}
	case DRIVER_PG:
		host := viper.GetString("database.pg.host")
		port := viper.GetString("database.pg.port")
		user := viper.GetString("database.pg.user")
		password := viper.GetString("database.pg.password")
		name := viper.GetString("database.pg.name")
		ssl := viper.GetString("database.pg.ssl")

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, name, port, ssl)
		log.Debug(dsn)
		db, err = gorm.Open("postgres", dsn)
		if err != nil {
			log.Fatal(fmt.Sprintf("Failed to connect postgres %s", err.Error()))
		} else {
			db.DB().SetMaxIdleConns(viper.GetInt("database.pg.pool.min"))
			db.DB().SetMaxOpenConns(viper.GetInt("database.pg.pool.max"))
			if gin.Mode() != gin.ReleaseMode {
				db.LogMode(true)
			}
		}
	default:
		log.Fatal("We do not support this kind of storage system yet!")
	}
	log.Info("Successfully connect to database")
}

// Shutdown - close database connection
func Shutdown() error {
	log.Info("Closing database's connections")
	return db.Close()
}

// GetDb - get a database connection
func GetDb() *gorm.DB {
	return db
}
