package inf

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"goutils/influx-db/influx"
	"testing"
)

func TestPoint(t *testing.T) {
	t.Skip()

	c = influx.NewClient()
	InitDB(DbName)

	Point(DemoName, 1)

	resp, err := c.Query(DbName, `select count(*) from `+DemoName)
	assert.NoError(t, err)
	fmt.Println(resp)
}

func BenchmarkPoint(b *testing.B) {
	b.Skip()

	b.StopTimer()

	c = influx.NewClient()
	InitDB(DbName)

	b.StartTimer()

	Point(DemoName, 1)
}
