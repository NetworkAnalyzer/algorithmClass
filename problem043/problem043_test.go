package problem043

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("trueのパターン", func(t *testing.T) {
		stack := NewStack()
		stack.Push(2)
		stack.Push(3)
		stack.Push(1)
		stack.Push(3)

		result, err := stack.Max()
		assert.Nil(t, err)
		assert.EqualValues(t, 3, result)

		result, err = stack.Pop()
		assert.Nil(t, err)
		assert.EqualValues(t, 3, result)

		result, err = stack.Max()
		assert.Nil(t, err)
		assert.EqualValues(t, 3, result)

		result, err = stack.Pop()
		assert.Nil(t, err)
		assert.EqualValues(t, 1, result)

		result, err = stack.Max()
		assert.Nil(t, err)
		assert.EqualValues(t, 3, result)

		result, err = stack.Pop()
		assert.Nil(t, err)
		assert.EqualValues(t, 3, result)

		result, err = stack.Max()
		assert.Nil(t, err)
		assert.EqualValues(t, 2, result)

		result, err = stack.Pop()
		assert.Nil(t, err)
		assert.EqualValues(t, 2, result)
	})

	t.Run("falseのパターン", func(t *testing.T) {
		stack := NewStack()

		_, err := stack.Max()
		assert.NotNil(t, err)

		_, err = stack.Pop()
		assert.NotNil(t, err)
	})
}
