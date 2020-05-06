package pointgenerate

import (
	"fmt"
	"testing"
)

func TestGenerate24(t *testing.T) {
	var res [][]int
	var testPrint func(arr []int)
	testPrint = func(arr []int) {
		res = append(res, arr)
	}
	Generate24(testPrint, SetMaxNumber(48))

	fmt.Println(len(res))
}
