package problem043

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestStack(t *testing.T) {
	t.Run("trueのパターン", func(t *testing.T) {
		stack := newStack()
		stack.push(2)
		stack.push(3)
		stack.push(1)
		stack.push(3)

		result, err := stack.max()
		assert.Nil(t, err)
		assert.EqualValues(t, 3, result)

		result, err = stack.pop()
		assert.Nil(t, err)
		assert.EqualValues(t, 3, result)

		result, err = stack.max()
		assert.Nil(t, err)
		assert.EqualValues(t, 3, result)

		result, err = stack.pop()
		assert.Nil(t, err)
		assert.EqualValues(t, 1, result)

		result, err = stack.max()
		assert.Nil(t, err)
		assert.EqualValues(t, 3, result)

		result, err = stack.pop()
		assert.Nil(t, err)
		assert.EqualValues(t, 3, result)

		result, err = stack.max()
		assert.Nil(t, err)
		assert.EqualValues(t, 2, result)

		result, err = stack.pop()
		assert.Nil(t, err)
		assert.EqualValues(t, 2, result)
	})

	t.Run("falseのパターン", func(t *testing.T) {
		stack := newStack()

		assert.EqualValues(t, stack.isEmpty(), true)

		_, err := stack.max()
		assert.NotNil(t, err)

		_, err = stack.pop()
		assert.NotNil(t, err)
	})
}
