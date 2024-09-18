package dao

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"wenzhi.com/gin-ranking/config"
	"wenzhi.com/gin-ranking/pkg/logger"
)

var (
	Db *gorm.DB
	err error
)

func init() {
	Db, err = gorm.Open("mysql", config.Mysqldb)
	if err!= nil {
		logger.Error(map[string]interface{}{"connect to mysql error: ": err.Error()})
	}
	if Db.Error!= nil {
		logger.Error(map[string]interface{}{"database error: ": Db.Error}) // Db.Error.Error()
	}

	// 连接池信息
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	Db.DB().SetConnMaxLifetime(time.Hour) // 10 minutes 10 * 60 * 1e9
}