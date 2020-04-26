package jwt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	token, err := Token(123)
	assert.NoError(t, err)
	//fmt.Println(token)

	id, err := Auth(token)
	assert.NoError(t, err)
	assert.Equal(t, 123, id)
}
