package custom

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEval32_no_function_assigned(t *testing.T) {
	gen := New(12345)

	require.Panics(t, func() {
		_ = gen.Eval32(1.0, 2.0, 3.0)
	}, "if the function is not assigned, it should panic")
}

func TestEval64_no_function_assigned(t *testing.T) {
	gen := New(12345)

	require.Panics(t, func() {
		_ = gen.Eval64(1.0, 2.0, 3.0)
	}, "if the function is not assigned, it should panic")
}

func TestSetEval32_function_returns_out_of_range(t *testing.T) {
	gen := New(12345)

	err := gen.SetEval32(func(seed int64, dim ...float32) float32 {
		return -2.0
	})

	require.Error(t, err, "if the function returns a value out of range, an error should be returned")
	assert.Contains(t, err.Error(), "the function must return a value between -1 and 1")
}

func TestSetEval64_function_returns_out_of_range(t *testing.T) {
	gen := New(12345)

	err := gen.SetEval64(func(seed int64, dim ...float64) float64 {
		return -2.0
	})

	require.Error(t, err, "if the function returns a value out of range, an error should be returned")
	assert.Contains(t, err.Error(), "the function must return a value between -1 and 1")
}
