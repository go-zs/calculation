package calculation24

import "fmt"

const (
	Add = iota + 1
	Subtract
	Multiply
	Divide

	diff = 0.00000001
)

var (
	operators = []int{Add, Subtract, Multiply, Divide}
)

type (
	Solution struct {
		Numbers   []int
		Operators []int
	}
)

func CalculatePoint(arr []int) []Solution {
	return calculatePoint(intToFloat32(arr))
}

func calculatePoint(arr []float32) []Solution {
	var backtrace func(numArr []float32, operatorArr []int, cur float32)
	var res []Solution
	numUsed := make([]bool, len(arr))
	backtrace = func(numsArr []float32, operatorArr []int, cur float32) {
		if len(numsArr) == 4 && CheckResult(cur) {
			res = append(res, Solution{
				Numbers:   CopyFloat32(numsArr),
				Operators: CopyInt(operatorArr),
			})
		}
		for i, n := range arr {
			if numUsed[i] {
				continue
			}
			numUsed[i] = true
			numsArr = append(numsArr, n)
			if len(numsArr) > 1 {
				for _, operator := range operators {
					operatorArr = append(operatorArr, operator)
					backtrace(numsArr, operatorArr, Calculation(cur, n, operator))
					operatorArr = operatorArr[:len(operatorArr)-1]
				}
			} else {
				backtrace(numsArr, operatorArr, n)
			}
			numUsed[i] = false
			numsArr = numsArr[:len(numsArr)-1]
		}
	}
	backtrace([]float32{}, []int{}, 1.0)

	return res
}

func CheckResult(r float32) bool {
	if c := r - 24.0; c < diff && c > -diff {
		return true
	}

	return false
}

func Calculation(n1, n2 float32, operator int) float32 {
	switch operator {
	case Add:
		return n1 + n2
	case Subtract:
		return n1 - n2
	case Multiply:
		return n1 * n2
	default:
		return n1 / n2
	}
}

func CopyFloat32(arr []float32) []int {
	res := make([]int, len(arr))
	for i := range arr {
		res[i] = int(arr[i])
	}
	return res
}

func CopyInt(arr []int) []int {
	res := make([]int, len(arr))
	copy(arr, res)
	return res
}

func PrintResult(res []Solution) []string {
	var stringRes []string
	for _, solution := range res {
		var i, j int
		var s string
		for i < len(solution.Numbers) {
			s += numberToString(solution.Numbers[i])
			i++
			if j < len(solution.Operators) {
				if j >= 1 && needBrackets(solution.Operators[j-1], solution.Operators[j]) {
					addBrackets(&s)
				}
				s += operatorToString(solution.Operators[j])
				j++
			}
		}
		stringRes = append(stringRes, s)
	}
	return stringRes
}

func numberToString(n int) string {
	return fmt.Sprintf("%d", n)
}

func operatorToString(o int) string {
	switch o {
	case Add:
		return " + "
	case Subtract:
		return " - "
	case Multiply:
		return " x "
	default:
		return " / "
	}
}

func needBrackets(o1, o2 int) bool {
	if (o2 == Multiply || o2 == Divide) && (o1 == Add || o1 == Subtract) {
		return true
	}

	return false
}

func addBrackets(s *string) {
	*s = fmt.Sprintf("(%s)", *s)
}

func intToFloat32(s []int) []float32 {
	floatArr := make([]float32, len(s))
	for i := range s {
		floatArr[i] = float32(s[i])
	}

	return floatArr
}
