package alephzero

// #cgo pkg-config: alephzero
// #include "common_adapter.h"
import "C"

import (
	"syscall"
	"unsafe"
)

func errorFrom(err C.errno_t) error {
	if err == 0 {
		return nil
	}
	return syscall.Errno(err)
}

///////////
// Alloc //
///////////

var (
	// TODO: make thread safe.
	allocRegistry = make(map[uintptr]func(C.size_t, *C.a0_buf_t))
	nextAllocId   uintptr
)

//export a0go_alloc
func a0go_alloc(id unsafe.Pointer, size C.size_t, out *C.a0_buf_t) {
	allocRegistry[uintptr(id)](size, out)
}

func registerAlloc(fn func(C.size_t, *C.a0_buf_t)) (id uintptr) {
	id = nextAllocId
	nextAllocId++
	allocRegistry[id] = fn
	return
}

func unregisterAlloc(id uintptr) {
	delete(allocRegistry, id)
}

//////////////
// Callback //
//////////////

var (
	// TODO: make thread safe.
	callbackRegistry = make(map[uintptr]func())
	nextCallbackId   uintptr
)

//export a0go_callback
func a0go_callback(id unsafe.Pointer) {
	callbackRegistry[uintptr(id)]()
}

func registerCallback(fn func()) (id uintptr) {
	id = nextCallbackId
	nextCallbackId++
	callbackRegistry[id] = fn
	return
}

func unregisterCallback(id uintptr) {
	delete(callbackRegistry, id)
}