package ex00

import (
	"errors"
	"unsafe"
)

func getElement(arr []int, idx int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("error: empty slice")
	} else if idx < 0 || idx >= len(arr) {
		return 0, errors.New("error: invalid index")
	}
	ptr := unsafe.Pointer(&arr[0])
	size := unsafe.Sizeof(0)
	return *(*int)(unsafe.Pointer(uintptr(ptr) + size*uintptr(idx))), nil
}
