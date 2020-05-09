package pointgenerate

import "github.com/go-zs/calculation/check"

const (
	defaultMaxNumber = 13
)

type (
	Option struct {
		maxNumber int
	}
	GenOption func(o *Option)
)

func SetMaxNumber(num int) GenOption {
	return func(o *Option) {
		o.maxNumber = num
	}
}

func Generate24(f func(arr []int), options ...GenOption) {
	option := Option{maxNumber: defaultMaxNumber}
	for _, o := range options {
		o(&option)
	}
	for i := 1; i <= option.maxNumber; i++ {
		for j := 1; j <= i; j++ {
			for m := 1; m <= j; m++ {
				for n := 1; n <= m; n++ {
					if check.Check24([]int{i, j, m, n}) {
						f([]int{i, j, m, n, getLevel(i, j, m, n)})
					}
				}
			}
		}
	}
}


func getLevel(i, j, m, n int) int {
	v := Max(i, j, m, n)
	switch  {
	case v <= 10:
		return 1
	case v <= 16:
		return 2
	case v <=24:
		return 3
	default:
		return 4
	}
}

func Max(intArr ...int) int {
	var m int
	for _, v := range intArr {
		if v > m {
			m = v
		}
	}

	return m
}