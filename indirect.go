//go:build !amd64
// +build !amd64

package indirect

import "unsafe"

func ToUnsafePtr(ptr uintptr) unsafe.Pointer {
	return unsafe.Pointer(ptr)
}
