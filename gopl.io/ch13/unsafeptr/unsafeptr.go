package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x struct {
		a bool
		b int16
		c []int
	}

	// 等价于 to pb := &x.b
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42
	fmt.Println(x.b) // "42"
}
