package format

import (
	"fmt"
	"time"
)

func Example() {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	fmt.Print(Any(x) + " ")
	fmt.Print(Any(d) + " ")
	fmt.Print(Any([]int64{x}) + " ")
	fmt.Print(Any([]time.Duration{d}))
}
