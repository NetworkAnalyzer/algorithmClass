package problem027

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("not error", func(t *testing.T) {
		result, err := Run("([])[]({})")
		assert.Nil(t, err)
		assert.EqualValues(t, true, result)

		result, err = Run("([)]")
		assert.Nil(t, err)
		assert.EqualValues(t, false, result)

		result, err = Run("((()")
		assert.Nil(t, err)
		assert.EqualValues(t, false, result)
	})

	t.Run("error", func(t *testing.T) {
		_, err := Run("(AA)[]")
		assert.NotNil(t, err)
	})
}
