package inf

import (
	"github.com/influxdata/influxdb1-client/v2"
	. "golib/lib/db/influx/client"
	"log"
	"os"
)

const (
	DbName   = "db"
	DemoName = "demo"
)

var c influxDb

type influxDb interface {
	Write(db string, name string, tags map[string]string, fields map[string]interface{}) error
	Query(db string, sql string) (*client.Response, error)
}

type mockClient struct{}

func (*mockClient) Write(string, string, map[string]string, map[string]interface{}) error { return nil }
func (*mockClient) Query(string, string) (*client.Response, error)                        { return nil, nil }

func init() {
	c = &mockClient{}
	if os.Getenv("INFLUX_ADDR") != "" {
		c = NewClient()
	}
	InitDB(DbName)
}

// if exists, do nothing and does not return an error
func InitDB(db string) {
	if _, err := c.Query("", "create database "+db); err != nil {
		log.Panic(err)
	}
}

func Point(name string, id uint) {
	if err := c.Write(DbName, name, nil, map[string]interface{}{"id": id}); err != nil {
		log.Println("write influx db err: " + err.Error())
	}
}
