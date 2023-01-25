package logs_parsing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPass(t *testing.T) {
	t.Parallel()

	assert.True(t, 1+1 == 2)

}

func TestFail(t *testing.T) {
	t.Parallel()

	assert.True(t, 1+1 == 3)

}
