package mathx_test

import (
	"testing"

	"github.com/alex-cos/mathx"
	"github.com/stretchr/testify/assert"
)

func TestMaximum(t *testing.T) {
	t.Parallel()
	assert.Equal(t, 3, mathx.Maximum(1, 2, 3))
	assert.Equal(t, 3.5, mathx.Maximum(1.0, 3.5, -2.0))
	assert.Equal(t, 7, mathx.Maximum(7))
	assert.Equal(t, 0, mathx.Maximum[int]())
	assert.Equal(t, 0.0, mathx.Maximum[float64]())
}

func TestMinimum(t *testing.T) {
	t.Parallel()
	assert.Equal(t, 1, mathx.Minimum(1, 3, 2))
	assert.Equal(t, -5.0, mathx.Minimum(1.0, -5.0, 2.0))
	assert.Equal(t, 7, mathx.Minimum(7))
	assert.Equal(t, 0, mathx.Minimum[int]())
	assert.Equal(t, 0.0, mathx.Minimum[float64]())
}

func TestMin(t *testing.T) {
	t.Parallel()
	assert.Equal(t, 1, mathx.Min([]int{1, 2, 3}))
	assert.Equal(t, -5.0, mathx.Min([]float64{1.0, -5.0, 2.0}))
	assert.Equal(t, 0, mathx.Min([]int{}))
	assert.Equal(t, 0.0, mathx.Min([]float64{}))
}

func TestMax(t *testing.T) {
	t.Parallel()
	assert.Equal(t, 3, mathx.Max([]int{1, 2, 3}))
	assert.Equal(t, 3.5, mathx.Max([]float64{1.0, -5.0, 3.5}))
	assert.Equal(t, 0, mathx.Max([]int{}))
	assert.Equal(t, 0.0, mathx.Max([]float64{}))
}

func TestSign(t *testing.T) {
	t.Parallel()
	assert.Equal(t, -1, mathx.Sign(-5))
	assert.Equal(t, 1, mathx.Sign(0))
	assert.Equal(t, 1, mathx.Sign(7))
	assert.Equal(t, -1.0, mathx.Sign(-3.14))
	assert.Equal(t, 1.0, mathx.Sign(2.71))
	assert.Equal(t, 1.0, mathx.Sign(0.0))
}
