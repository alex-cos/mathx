package mathx_test

import (
	"testing"

	"github.com/alex-cos/mathx"
	"github.com/stretchr/testify/assert"
)

func TestAverage(t *testing.T) {
	t.Parallel()
	assert.Equal(t, 2.0, mathx.Average[int, float64]([]int{1, 2, 3}))
	assert.Equal(t, 0.0, mathx.Average[int, float64]([]int{}))
	assert.Equal(t, 2.0, mathx.Average[float64, float64]([]float64{1.0, 2.0, 3.0}))
}
