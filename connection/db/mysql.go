package db

import (
	"fmt"
	"hello_blockchain/config"
	"time"

	"github.com/rotisserie/eris"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type _mysql struct {
	db *gorm.DB // mysql连线
}

func NewMysql() *_mysql {
	return &_mysql{}
}

func (c *_mysql) Init() {
	dsn := config.MysqlConn.Username + ":" + config.MysqlConn.Password +
		"@tcp(" + config.MysqlConn.Host + ":" + config.MysqlConn.Port + ")/" + config.MysqlConn.Database +
		"?charset=" + config.MysqlConn.Charset + "&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info), // debug 用, 更新生产前需要移除
		SkipDefaultTransaction: true,                                // gorm 在 create/update/delete 会自动用 transaction 包起来, 效能容易不好
		DisableAutomaticPing:   true,
	})
	if err != nil {
		fmt.Println("dsn:", dsn)
		panic(eris.Wrap(err, "mysql init error"))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(eris.Wrap(err, "mysql get sqlDB error"))
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	err = sqlDB.Ping()
	if err != nil {
		panic(eris.Wrap(err, "mysql ping error"))
	} else {
		c.db = db
	}
}

func (c *_mysql) Close() {

}
