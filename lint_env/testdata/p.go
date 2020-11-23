package testdata

import (
	"fmt"
	"unsafe"
)

func p() {
	const size = unsafe.Sizeof(uint32(0))
	fmt.Println(size)
}
