package client

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient(t *testing.T) {
	t.Skip()

	db := "test"
	name := "point"
	tag := map[string]string{} //map[string]string{"id": "1"}
	fields := map[string]interface{}{"field": "text"}
	query := `SELECT * FROM ` + name
	newDb := `create database ` + db
	delDb := `drop database ` + db

	cli := NewClient()

	_, err := cli.Query(db, newDb)
	assert.NoError(t, err)

	err = cli.Write(db, name, tag, fields)
	assert.NoError(t, err)

	resp, err := cli.Query(db, query)
	assert.NoError(t, err)
	fmt.Println(resp)

	_, err = cli.Query(db, delDb)
	assert.NoError(t, err)

	err = cli.Close()
	assert.NoError(t, err)
}
