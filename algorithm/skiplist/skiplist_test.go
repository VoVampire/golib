package skiplist

import (
	//"github.com/sanity-io/litter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZSkipList(t *testing.T) {
	zsl := ZslCreate()
	zsl.ZslInsert(100, "1")
	zsl.ZslInsert(200, "2")
	zsl.ZslInsert(300, "3")
	x := zsl.ZslQuery(200, "2")
	assert.NotEmpty(t, x)

	assert.Equal(t, 1, zsl.ZslDelete(200, "2"))
	assert.Equal(t, 0, zsl.ZslDelete(200, "2"))

	x = zsl.ZslQuery(200, "2")
	assert.Empty(t, x)
}
