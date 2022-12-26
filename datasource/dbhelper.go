package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lottery/conf"
)

type MysqlConn struct {
	DB *gorm.DB
}

func New() *MysqlConn {
	return &MysqlConn{
		DB: newDb(),
	}
}

func newDb() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DBMaster.User, conf.DBMaster.Port, conf.DBMaster.Host, conf.DBMaster.Port, conf.DBMaster.Databases)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(any("连接数据库失败, error=" + err.Error()))
	}

	sqlDB, _ := db.DB()

	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	return db
}
