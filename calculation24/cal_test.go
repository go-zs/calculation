package calculation24

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckResult(t *testing.T) {
	assert.True(t, CheckResult(12.0 * 2.0))
	assert.True(t, CheckResult(48.0 / 2.0))
	assert.True(t, CheckResult(3.0 * 8.0))
	assert.True(t, CheckResult(6.0 * 8.0 / 2.0))
	assert.True(t, CheckResult(23.0 + 1.0))
	assert.True(t, CheckResult(22.5 + 1.5))
	assert.False(t, CheckResult(48.01))
	assert.False(t, CheckResult(12.01))
	assert.False(t, CheckResult(13.01))
}

func TestCalculation(t *testing.T) {
	assert.Equal(t, float32(2.0 + 2.0) , Calculation(2.0, 2.0, Add))
	assert.Equal(t, float32(2.0 - 2.0) , Calculation(2.0, 2.0, Subtract))
	assert.Equal(t, float32(2.0 * 2.0) , Calculation(2.0, 2.0, Multiply))
	assert.Equal(t, float32(2.0 / 2.0) , Calculation(2.0, 2.0, Divide))
}

func TestCalPoint(t *testing.T) {
	result := CalculatePoint([]int{5, 2, 3, 4})
	arr := PrintResult(result)
	for _, s := range arr {
		fmt.Println(s)
	}
}