package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLabel(t *testing.T) {
	t.Run("test that label is correct format", func(t *testing.T) {
		assert.Equal(t, []byte("foo/bar"), NewLabel("foo", "bar"))
	})
}
