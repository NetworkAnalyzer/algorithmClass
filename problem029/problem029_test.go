package problem029

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestEncode(t *testing.T) {
	before := "AAAABBBCCDAA"
	after := "4A3B2C1D2A"

	t.Run("trueのパターン", func(t *testing.T) {
		result, err := encode(before)

		assert.Nil(t, err)
		assert.EqualValues(t, after, result)
	})

	t.Run("falseのパターン", func(t *testing.T) {
		_, err := encode(after)

		assert.NotNil(t, err)
	})
}

func TestDecode(t *testing.T) {
	before := "4A3B2C1D2A"
	after := "AAAABBBCCDAA"

	t.Run("trueのパターン", func(t *testing.T) {
		result, err := decode(before)

		assert.Nil(t, err)
		assert.EqualValues(t, after, result)
	})

	t.Run("falseのパターン", func(t *testing.T) {
		_, err := decode(after)
		assert.NotNil(t, err)

		_, err = decode(before + "AA")
		assert.NotNil(t, err)
	})
}
