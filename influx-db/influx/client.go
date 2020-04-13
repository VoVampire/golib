package influx

import (
	client "github.com/influxdata/influxdb1-client/v2"
	"log"
	"time"
)

type Client struct {
	cli client.Client
}

func NewClient() *Client {
	con, err := client.NewHTTPClient(httpConfig())
	if err != nil {
		log.Fatal(err)
	}

	if _, _, err := con.Ping(time.Second); err != nil {
		log.Fatal(err)
	}

	log.Println("Connection", con)
	return &Client{cli: con}
}

func (i *Client) Query(db string, sql string) (*client.Response, error) {
	resp, err := i.cli.Query(client.NewQuery(sql, db, ""))
	if err != nil {
		return nil, err
	}

	return resp, resp.Error()
}

func (i *Client) Write(db string, name string, tags map[string]string, fields map[string]interface{}) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{Database: db})
	if err != nil {
		return err
	}

	p, err := client.NewPoint(name, tags, fields, time.Now())
	if err != nil {
		return err
	}

	bp.AddPoint(p)

	return i.cli.Write(bp)
}

func (i *Client) Close() error {
	return i.cli.Close()
}
