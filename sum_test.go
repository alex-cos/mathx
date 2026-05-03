package mathx_test

import (
	"testing"

	"github.com/alex-cos/mathx"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Parallel()
	assert.Equal(t, 6, mathx.Sum([]int{1, 2, 3}))
	assert.Equal(t, 7.5, mathx.Sum([]float64{2.5, 5.0}))
	assert.Equal(t, 0, mathx.Sum([]int{}))
	assert.Equal(t, 0.0, mathx.Sum([]float64{}))
}
