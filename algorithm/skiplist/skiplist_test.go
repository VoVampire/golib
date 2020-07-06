package skiplist

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZSkipList_Insert(t *testing.T) {
	zsl := Create()
	zsl.Insert(100, "1")
	zsl.Insert(200, "2")
	bts, err := json.Marshal(zsl)
	assert.NoError(t, err)
	fmt.Println(string(bts))
	//fmt.Println(jsoniter.MarshalToString(zsl))
}
