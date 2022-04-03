package indirect

import "unsafe"

//go:noescape
func indUP(_ uintptr) unsafe.Pointer
