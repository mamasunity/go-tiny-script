package prime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrime(t *testing.T) {
	assert.Equal(t, false, IsPrime(1))
	assert.Equal(t, true, IsPrime(5))
	assert.Equal(t, false, IsPrime(8))
	assert.Equal(t, false, IsPrime(39))
	assert.Equal(t, true, IsPrime(41))
}
