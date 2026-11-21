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

func TestAverage_NoIntegerOverflow(t *testing.T) {
	t.Parallel()
	assert.Equal(t, 100.0, mathx.Average[int8, float64]([]int8{100, 100, 100}))
	assert.Equal(t, -100.0, mathx.Average[int8, float64]([]int8{-100, -100, -100}))
	assert.Equal(t, 30000.0, mathx.Average[int16, float64]([]int16{30000, 30000, 30000}))
	assert.Equal(t, 2000000000.0, mathx.Average[int32, float64]([]int32{2000000000, 2000000000}))
}

func TestAverage_Float32Output(t *testing.T) {
	t.Parallel()
	assert.Equal(t, float32(2.5), mathx.Average[int, float32]([]int{2, 3}))
}
