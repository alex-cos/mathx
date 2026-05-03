package mathx_test

import (
	"testing"

	"github.com/alex-cos/mathx"
	"github.com/stretchr/testify/assert"
)

func TestVariance(t *testing.T) {
	t.Parallel()
	assert.Equal(t, 1.0, mathx.Variance[float64, float64]([]float64{1.0, 2.0, 3.0}))
	assert.Equal(t, 1.0, mathx.Variance[int, float64]([]int{1, 2, 3}))
	assert.Equal(t, 0.0, mathx.Variance[float64, float64]([]float64{5.0}))
	assert.Equal(t, 0.0, mathx.Variance[float64, float64]([]float64{}))
}

func TestStdDev(t *testing.T) {
	t.Parallel()
	assert.Equal(t, 1.5811388300841898, mathx.StdDev[float64, float64]([]float64{1, 2, 3, 4, 5}))
	assert.Equal(t, 2.138089935299395, mathx.StdDev[float64, float64]([]float64{2, 4, 4, 4, 5, 5, 7, 9}))
	assert.Equal(t, 0.0, mathx.StdDev[float64, float64]([]float64{}))
	assert.Equal(t, 0.0, mathx.StdDev[float64, float64]([]float64{1}))
}
