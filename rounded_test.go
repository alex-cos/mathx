package mathx_test

import (
	"testing"

	"github.com/alex-cos/mathx"
	"github.com/stretchr/testify/assert"
)

func TestRounded(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input     float64
		precision int
		increment float64
		upper     bool
		expected  float64
	}{
		{12.18978, 2, 0.05, false, 12.2},
		{12.18978, 2, 0.05, true, 12.20},
		{112.3467986, 4, 0.0002, false, 112.3468},
		{112.3467986, 4, 0.0002, true, 112.3468},
		{34.79, 2, 0.05, true, 34.8},
		{201.3489, 2, 0.0001, true, 201.34},
	}
	for i, test := range tests {
		output := mathx.Rounded(test.input, test.precision, test.increment, test.upper)
		assert.Equalf(t, test.expected, output, "test sample #%d", i)
	}
}
