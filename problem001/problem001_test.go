package problem001

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("trueのパターン", func(t *testing.T) {
		actual := Run([]int{10, 13, 5, 7}, 17)

		assert.EqualValues(t, true, actual)
	})

	t.Run("falseのパターン", func(t *testing.T) {
		actual := Run([]int{10, 13, 5, 8}, 17)

		assert.EqualValues(t, false, actual)
	})
}
