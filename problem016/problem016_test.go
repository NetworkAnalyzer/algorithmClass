package problem016

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecord(t *testing.T) {
	expected := []log{
		{
			orderID: 1,
		},
		{
			orderID: 2,
		},
	}

	t.Run("trueのパターン", func(t *testing.T) {
		var err error

		err = record(1)
		err = record(2)

		assert.Nil(t, err)
		assert.EqualValues(t, expected, logs)
	})

	t.Run("already existsのパターン", func(t *testing.T) {
		err := record(1)

		assert.NotNil(t, err)
	})
}

func TestGetLast(t *testing.T) {
	for _, v := range []int{1, 2, 3, 4, 5} {
		record(v)
	}

	last1 := []log{
		{
			orderID: 5,
		},
	}
	assert.EqualValues(t, last1, getLast(1))

	last2 := []log{
		{
			orderID: 4,
		},
		{
			orderID: 5,
		},
	}
	assert.EqualValues(t, last2, getLast(2))
}
