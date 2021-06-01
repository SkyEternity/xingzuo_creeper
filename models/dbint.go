package models

import (
	"log"
	"time"

	"github.com/mattn/go-colorable"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	err error
)

var newLogger = logger.New(
	// io.writer同样使用colorable
	log.New(colorable.NewColorableStdout(), "\r\n", log.LstdFlags),
	logger.Config{
		SlowThreshold: time.Second, // 慢 SQL 阈值
		LogLevel:      logger.Info, // Log level
		Colorful:      true,        // 开启彩色打印
	},
)

func init() {
	dbName := "root"
	dbPwd := "123456"
	dbUrl := "127.0.0.1:3306"
	dbDatabase := "creeper"
	dsn := dbName + ":" + dbPwd + "@tcp(" + dbUrl + ")/" + dbDatabase + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: newLogger, //开启sql日志
	})
	DB.Debug()
}
