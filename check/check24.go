package check

import "github.com/go-zs/calculation/calculation24"

func Check24(arr []int) bool {
	res := calculation24.CalculatePoint(arr)

	return len(res) > 0
}
