package indirect

import "unsafe"

func ToUnsafePtr(ptr uintptr) unsafe.Pointer {
	return indUP(ptr)
}
