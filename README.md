# Mathx

[![Go Version](https://img.shields.io/badge/Go-1.25%2B-blue)](https://go.dev/)
[![Test Status](https://github.com/alex-cos/mathx/actions/workflows/test.yml/badge.svg)](https://github.com/alex-cos/mathx/actions/workflows/test.yml)
[![Lint Status](https://github.com/alex-cos/mathx/actions/workflows/lint.yml/badge.svg)](https://github.com/alex-cos/mathx/actions/workflows/lint.yml)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/alex-cos/mathx)](https://goreportcard.com/report/github.com/alex-cos/mathx)

A lightweight, generic statistics library for Go

This module provides a small collection of statistical functions implemented with Go generics. It is designed to operate on numeric data without forcing a specific underlying type, enabling reuse across different numeric types (integers and floating-point numbers).

## Overview

- Core idea: expose common statistical operations as generic functions that work with any numeric type implementing the library's constraints.
- Key types are defined in types.go:
  - Integer: int, int8, int16, int32, int64
  - Float: float32, float64
  - Number: Integer | Float (union of the above)

## Statistical Functions

### Maximum

```go
Maximum[N Number](nums ...N) N
```

Returns the maximum value among the given variadic numbers.

- **Edge cases**: if no arguments are provided, returns the zero value of N (e.g. 0 or 0.0).
- **Example**: `Maximum[int](1, 5, 3)` → `5`

---

### Minimum

```go
Minimum[N Number](nums ...N) N
```

Returns the minimum value among the given variadic numbers.

- **Edge cases**: if no arguments are provided, returns the zero value of N.
- **Example**: `Minimum[int](1, -2, 7)` → `-2`

---

### Sign

```go
Sign[N Number](v N) N
```

Returns -1 if v is negative; otherwise returns +1 (including 0).

- **Example**: `Sign(-3)` → `-1`, `Sign(0)` → `1`, `Sign(4)` → `1`

---

### Sum

```go
Sum[N Number](values []N) N
```

Sums all values in the provided slice and returns the result.

- **Edge cases**: empty slice → 0 (zero value of N)
- **Example**: `Sum([]int{1,2,3})` → `6`

---

### Min

```go
Min[N Number](values []N) N
```

Slice-based equivalent of Minimum. Returns the minimum value in the slice.

- **Behavior on empty input**: returns the zero value of N.
- **Example**: `Min([]int{3, 1, 2})` → `1`

---

### Max

```go
Max[N Number](values []N) N
```

Slice-based equivalent of Maximum. Returns the maximum value in the slice.

- **Behavior on empty input**: returns the zero value of N.
- **Example**: `Max([]int{3, 1, 2})` → `3`

---

### Average

```go
Average[N Number, F Float](values []N) F
```

Computes the arithmetic mean of the values and returns it as type F.

- **Edge cases**: empty slice → 0.0
- **Example**: `Average[int, float64]([]int{1,2,3})` → `2.0`

---

### Variance

```go
Variance[N Number, F Float](values []N) F
```

Computes the sample variance (denominator n-1) of the given values and returns type F.

- **Edge cases**: 0 or 1 element → 0.0
- **Implemented as**: `(sum((x - mean)^2)) / (n-1)`
- **Example**: `Variance[float64]([]float64{1,2,3})` → `1.0`

---

### StdDev

```go
StdDev[N Number, F Float](values []N) F
```

Returns the sample standard deviation of the values.

- **Formula**: `sqrt(Variance(values))`
- **Example**: `StdDev[float64]([]float64{1,2,3,4,5})` → `1.581...`

---

### Covariance

```go
Covariance[N Number, F Float](x, y []N) F
```

Computes the sample covariance between two variables.

- **Formula**: `Σ((xi - x̄)(yi - ȳ)) / (n-1)`
- **Edge cases**:
  - slices of different lengths → 0
  - 0 or 1 element → 0
- **Example**: `Covariance([]float64{1,2,3,4,5}, []float64{1,2,3,4,5})` → `2.5`

---

### Correlation

```go
Correlation[N Number, F Float](x, y []N) F
```

Computes the Pearson correlation coefficient between x and y.

- **Formula**: `Covariance(x, y) / (StdDev(x) * StdDev(y))`
- **Range**: -1 to 1 (1 = perfect positive correlation, -1 = perfect negative)
- **Example**: `Correlation([]float64{1,2,3,4,5}, []float64{1,2,3,4,5})` → `1`

---

### Rounded

```go
Rounded[F Float](value F, precision int, increment F, upper bool) F
```

Rounds a value at a given decimal precision to a specified increment.

- **precision**: controls the decimal places (via p = 10^precision).
- **increment**: if > 0, rounding targets multiples of the increment.
- **upper**: if true, rounding pushes the result upward to the next increment when applicable.
- **Example**: `Rounded(12.18978, 2, 0.05, false)` → `12.20`

---

### Percentile

```go
Percentile[F Float](values []F, p F) F
```

Computes the p-th percentile of the values (0 <= p <= 1) using linear interpolation.

- **Edge cases**:
  - n==0 → 0
  - n==1 → the single value
  - p outside [0,1] clamps to ends
- **Example**: `Percentile([]float64{1,2,3,4}, 0.5)` → `2.5`

---

### Median

```go
Median[F Float](values []F) F
```

Computes the median of the values (50th percentile).

- Uses `Percentile` with `p = 0.5`.
- **Edge cases**:
  - empty slice → 0
  - single value → the value
  - even n → average of two middle values
- **Example**: `Median([]float64{1, 2, 3})` → `2`
- **Example**: `Median([]float64{1, 2, 3, 4})` → `2.5`

---

### Mode

```go
Mode[F Float](values []F) F
```

Computes the mode (most frequent value) of the values.

- Returns the first occurring mode if there are ties.
- **Edge cases**:
  - empty slice → 0
  - single value → the value
  - all values unique → returns the first value
- **Example**: `Mode([]float64{1, 2, 2, 3})` → `2`

---

### MinMaxNormalize

```go
MinMaxNormalize[F Float](values []F) []F
```

Normalizes values to [0, 1] range using min-max scaling.

- **Formula**: `(value - min) / (max - min)`
- **Edge cases**:
  - empty slice → empty slice
  - single value → `[1]`
  - all values equal → all values set to `1`
- **Example**: `MinMaxNormalize([]float64{1, 2, 3})` → `[0, 0.5, 1]`

---

### SumNormalize

```go
SumNormalize[F Float](values []F) []F
```

Normalizes values so their sum equals 1.

- **Formula**: `value / sum(values)`
- **Edge cases**:
  - empty slice → empty slice
  - all values zero → slice of zeros
- **Example**: `SumNormalize([]float64{1, 2, 3})` → `[0.166..., 0.333..., 0.5]`

---

### ComputeDistribution

```go
ComputeDistribution[F Float, I Integer](values []F, interval, precision int) []DistributionInterval[F, I]
```

Builds a distribution (histogram-like) of the input values across 'interval' bins.

Returns a slice of DistributionInterval with each interval's lower bound (Inf), upper bound (Sup) and a Count of values in that bin.

- Uses "nice" numbers for intervals: step sizes are rounded to 1, 2, 5, 10, 20, 50, 100, etc. for better readability.
- Boundaries are aligned to multiples of this step for clean intervals.
- **Example**: values in [-10, 11.2] with 4 intervals → `[-10, 0], [0, 10], [10, 20]` (step = 10)

---

## Math Functions

### Pow

```go
Pow[F Float](value F, exponent F) F
```

Returns the value raised to the specified exponent.

- **Formula**: `value^exponent`
- **Example**: `Pow(4, 2)` → `16`
- **Example**: `Pow(27, 1.0/3.0)` → `3`

---

### PowAll

```go
PowAll[F Float](values []F, exponent F) []F
```

Returns each value raised to the specified exponent.

- **Example**: `PowAll([]float64{1, 2, 3}, 2)` → `[1, 4, 9]`

---

### Sqrt

```go
Sqrt[F Float](value F) F
```

Returns the square root of the value.

- **Example**: `Sqrt(16)` → `4`
- **Example**: `Sqrt(2)` → `1.414...`

---

### SqrtAll

```go
SqrtAll[F Float](values []F) []F
```

Returns the square root of each value.

- **Example**: `SqrtAll([]float64{4, 9, 16})` → `[2, 3, 4]`

---

### Abs

```go
Abs[F Float](value F) F
```

Returns the absolute value.

- **Example**: `Abs(-5)` → `5`
- **Example**: `Abs(3.14)` → `3.14`

---

### AbsAll

```go
AbsAll[F Float](values []F) []F
```

Returns the absolute value of each value.

- **Example**: `AbsAll([]float64{-1, 2, -3})` → `[1, 2, 3]`

---

## Finance Functions

### DailyReturn

```go
DailyReturn[F Float](values []F) []F
```

Computes daily returns from price values.

- **Formula**: `(Pt+1 / Pt) - 1`
- **Example**: `DailyReturn([]float64{100, 110, 121})` → `[0.1, 0.1]`

---

### Volatility

```go
Volatility[F Float](values []F) F
```

Computes the volatility (standard deviation) of the returns.

- **Formula**: `sqrt(Variance(returns))`
- **Example**: `Volatility([]float64{0.01, -0.02, 0.03})` → `0.0255...`

---

### AnnualizedReturn

```go
AnnualizedReturn[F Float](values []F, nbDays F) F
```

Computes the annualized return from returns.

- **Formula**: `(Π(1 + ri))^(nbDays/n) - 1`
- **Example**: `AnnualizedReturn([]float64{0.01, -0.02, 0.03}, 252)` → annual return

---

### AnnualizedVolatility

```go
AnnualizedVolatility[F Float](values []F, nbDays F) F
```

Computes the annualized volatility.

- **Formula**: `Volatility(returns) * sqrt(nbDays)`
- **Example**: `AnnualizedVolatility([]float64{0.01, -0.02}, 252)` → `0.223...`

---

### ComputeDrawdownFromDR

```go
ComputeDrawdownFromDR[N Number](dailyReturns []N) []N
```

Computes drawdown from daily returns.

- **Example**: `ComputeDrawdownFromDR([]float64{-0.01, 0.02, -0.03})` → drawdowns

---

### ComputeCumulPerfFromDR

```go
ComputeCumulPerfFromDR[N Number](dailyReturns []N) []N
```

Computes cumulative performance from daily returns.

- **Example**: `ComputeCumulPerfFromDR([]float64{0.01, -0.02})` → cumulated performance

---

### MaxDrawdown

```go
MaxDrawdown[N Number](dailyReturns []N) N
```

Returns the maximum drawdown from daily returns.

- **Example**: `MaxDrawdown([]float64{-0.01, 0.02, -0.05})` → `-0.0406...`

---

## Example usage

```go
package main

import (
    "fmt"
    "github.com/alex-cos/mathx"
)

func main() {
    // Maximum / Minimum using generics
    fmt.Println(mathx.Maximum(1, 4, 2)) // 4
    fmt.Println(mathx.Minimum(1, -2, 3)) // -2

    // Other statistics
    nums := []int{1, 2, 3, 4, 5}
    fmt.Println(mathx.Average[int, float64](nums)) // 3.0
    fmt.Println(mathx.Variance[float64, float64]([]float64{1, 2, 3, 4, 5})) // 2.5
}
```
