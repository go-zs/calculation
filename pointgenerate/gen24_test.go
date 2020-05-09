package pointgenerate

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerate24(t *testing.T) {
	var res [][]int
	var testPrint = func(arr []int) {
		res = append(res, arr)
	}
	Generate24(testPrint, SetMaxNumber(24))
	fmt.Println(res)
}

func TestMax(t *testing.T) {
	arr := []int{1,2,3,4,5}
	assert.Equal(t, 5, Max(arr...))
}