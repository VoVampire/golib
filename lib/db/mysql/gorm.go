package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"strings"
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

func NewMysql(c *Conf) *gorm.DB {
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", c.User, c.Pwd, c.Addr, c.Dbname, c.Charset)
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(c.MaxIdle)
	db.DB().SetMaxOpenConns(c.MaxOpen)
	db.DB().SetConnMaxLifetime(c.MaxLifetime)
	return db
}

func NewExample() *gorm.DB {
	user, _ := ioutil.ReadFile("/mysql.usr")
	usr := strings.TrimSpace(string(user))

	password, _ := ioutil.ReadFile("/mysql.pwd")
	pwd := strings.TrimSpace(string(password))

	conn := fmt.Sprintf("%s:%s@/gadmin_test?charset=utf8&parseTime=True&loc=Local", usr, pwd)
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		panic(err)
	}
	return db
}
