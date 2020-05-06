package check

import "github.com/go-zs/cal/calculation24"

func Check24(arr []int) bool {
	res := calculation24.CalculatePoint(arr)

	if len(res) > 0 {
		return true
	}

	return false
}



