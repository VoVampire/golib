package gorm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type Conf struct {
	User        string
	Pwd         string
	Addr        string // "localhost:3306"
	Dbname      string
	Charset     string // "utf8","utf8mb4"
	MaxIdle     int
	MaxOpen     int
	MaxLifetime time.Duration
}

func NewConn(c *Conf) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", c.User, c.Pwd, c.Addr, c.Dbname)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(c.MaxIdle)
	db.DB().SetMaxOpenConns(c.MaxOpen)
	db.DB().SetConnMaxLifetime(c.MaxLifetime)
	return db
}

func CreateDB(conf *Conf) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/sys?charset=utf8&parseTime=True&loc=Local", conf.User, conf.Pwd, conf.Addr)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("连接数据库出错:" + err.Error())
	}

	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARSET utf8 COLLATE utf8_general_ci;", conf.Dbname)
	err = db.Exec(sql).Error
	if err != nil {
		log.Println("数据库创建失败:" + err.Error())
		return
	}
	log.Println("数据库创建成功:" + conf.Dbname)
}
