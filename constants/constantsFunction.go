package constants

import (
	"fmt"
	"math"
)

const s string = "constant"

func ConstantsFunction() {
	fmt.Println(s)

	const n = 50000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
