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

	logger := newLogger()

	t.Run("trueのパターン", func(t *testing.T) {
		var err error

		err = logger.record(1)
		err = logger.record(2)

		assert.Nil(t, err)
		assert.EqualValues(t, expected, logger.log)
	})

	t.Run("already existsのパターン", func(t *testing.T) {
		err := logger.record(1)

		assert.NotNil(t, err)
	})
}

func TestGetLast(t *testing.T) {
	logger := newLogger()

	for _, v := range []int{1, 2, 3, 4, 5} {
		logger.record(v)
	}

	last1 := []log{
		{
			orderID: 5,
		},
	}
	assert.EqualValues(t, last1, logger.getLast(1))

	last2 := []log{
		{
			orderID: 4,
		},
		{
			orderID: 5,
		},
	}
	assert.EqualValues(t, last2, logger.getLast(2))
}
