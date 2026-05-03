package mathx_test

import (
	"testing"

	"github.com/alex-cos/mathx"
)

func BenchmarkPow(b *testing.B) {
	for b.Loop() {
		mathx.Pow(2.0, 10.0)
	}
}

func BenchmarkSqrt(b *testing.B) {
	for b.Loop() {
		mathx.Sqrt(1024.0)
	}
}

func BenchmarkAbs(b *testing.B) {
	for b.Loop() {
		mathx.Abs(-1234.5678)
	}
}

func BenchmarkSum(b *testing.B) {
	data := makeFloatSlice()
	for b.Loop() {
		mathx.Sum(data)
	}
}

func BenchmarkAverage(b *testing.B) {
	data := makeIntSlice()
	for b.Loop() {
		mathx.Average[int, float64](data)
	}
}

func BenchmarkVariance(b *testing.B) {
	data := makeFloatSlice()
	for b.Loop() {
		mathx.Variance[float64, float64](data)
	}
}

func BenchmarkStdDev(b *testing.B) {
	data := makeFloatSlice()
	for b.Loop() {
		mathx.StdDev[float64, float64](data)
	}
}

func BenchmarkMedian(b *testing.B) {
	data := makeFloatSlice()
	for b.Loop() {
		mathx.Median(data)
	}
}

func BenchmarkMode(b *testing.B) {
	data := makeFloatSlice()
	for b.Loop() {
		mathx.Mode(data)
	}
}

func BenchmarkCorrelation(b *testing.B) {
	x := makeFloatSlice()
	y := makeFloatSlice()
	for b.Loop() {
		mathx.Correlation[float64, float64](x, y)
	}
}

func BenchmarkComputeDistribution(b *testing.B) {
	data := makeFloatSlice()
	for b.Loop() {
		mathx.ComputeDistribution[float64, int](data, 10, 2)
	}
}

func BenchmarkMinMaxNormalize(b *testing.B) {
	data := makeFloatSlice()
	for b.Loop() {
		mathx.MinMaxNormalize(data)
	}
}

func BenchmarkDailyReturn(b *testing.B) {
	data := makeFloatSlice()
	for b.Loop() {
		mathx.DailyReturn(data)
	}
}

func makeFloatSlice() []float64 {
	data := make([]float64, 1000)
	for i := range 1000 {
		data[i] = float64(i)
	}
	return data
}

func makeIntSlice() []int {
	data := make([]int, 1000)
	for i := range 1000 {
		data[i] = i
	}
	return data
}
