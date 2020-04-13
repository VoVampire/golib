package influx

import (
	client "github.com/influxdata/influxdb1-client/v2"
	"os"
)

func httpConfig() client.HTTPConfig {
	cfg := client.HTTPConfig{
		Addr:     os.Getenv("INFLUX_ADDR"),
		Username: os.Getenv("INFLUX_USER"),
		Password: os.Getenv("INFLUX_PWD"),
	}
	if cfg.Addr == "" {
		cfg.Addr = "http://localhost:8086"
	}
	return cfg
}
