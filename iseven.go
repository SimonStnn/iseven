package iseven

import "github.com/SimonStnn/isodd"

func IsEven(i int) bool {
	return !isodd.IsOdd(i)
}