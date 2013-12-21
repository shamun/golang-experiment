package main

import (
	"fmt"
	"unsafe"
)

func InspectSlice(slice []string) {
	address := unsafe.Pointer(&slice)

	lenAddr := uintptr(address) + uintptr(8)
	capAddr := uintptr(address) + uintptr(16)

	lenPtr := (*int)(unsafe.Pointer(lenAddr))
	capPtr := (*int)(unsafe.Pointer(capAddr))

	addPtr := (*[8]string)(unsafe.Pointer(*(*uintptr)(address)))

	fmt.Printf("Slice Addr[%p] Len Addr[0x%x] Cap Addr[0x%x]\n",
		address,
		lenAddr,
		capAddr)

	fmt.Printf("Slice Length[%d] Cap[%d]\n",
		*lenPtr,
		*capPtr)

	for index := 0; index < *lenPtr; index++ {
		fmt.Printf("[%d] %p %s\n",
			index,
			&(*addPtr)[index],
			(*addPtr)[index])
	}

	fmt.Printf("\n\n")
}

func main() {
	source := []string{"Apple", "Orange", "plum", "Banana", "grape"}
	InspectSlice(source)

	takeOne := source[2:3]
	InspectSlice(takeOne)

	takeOneCapOne := source[2:3:3]
	InspectSlice(takeOneCapOne)

	takeOneCapOne = append(takeOneCapOne, "kiwi")
	InspectSlice(takeOneCapOne)

	//danger before
	takeOne = append(takeOne, "rock")
	InspectSlice(takeOne)
	InspectSlice(source)
}
