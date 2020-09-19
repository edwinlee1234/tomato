package model

import (
	"fmt"
	"time"

	"tomato/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBConn DBConn
var DBConn *gorm.DB

// Init Init
func Init() {
	DBConn = connDB(
		config.Val.DBUser,
		config.Val.DBPass,
		config.Val.DBHost,
		config.Val.DBPort,
		config.Val.DBDatabase,
		config.Val.DBMaxLifeTime,
		config.Val.DBMaxConn,
		config.Val.DBIdleConn)
}

func connDB(user, pass, host, port, database string, lifeTime, maxCon, idle int) *gorm.DB {
	addr := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=true",
		user,
		pass,
		host,
		port,
		database,
	)

	db, err := gorm.Open(mysql.Open(addr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetConnMaxLifetime(time.Duration(lifeTime) * time.Second) // 每條連線的存活時間
	sqlDB.SetMaxOpenConns(maxCon)                                   // 最大連線數
	sqlDB.SetMaxIdleConns(idle)                                     // 最大閒置連線數

	return db
}
