package indirect

import "unsafe"

func ToUnsafePtr(ptr uintptr) unsafe.Pointer {
	return indUP(ptr)
}

//go:noescape
func indUP(_ uintptr) unsafe.Pointer
