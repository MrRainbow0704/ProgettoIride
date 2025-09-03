//go:build windows

// Il pacchetto video espone le funzioni dell'SDK in C++ di
// ActiveSilicon per la cattura di video.
package video

/*
#cgo pkg-config: aravis-0.8 glib-2.0 libpng
#include "stdlib.h"
#include "video.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type Camera struct {
	ptr *C.ArvCamera
}

type Buffer struct {
	ptr *C.ArvBuffer
}

func NewCamera() (*Camera, error) {
	cam := &C.ArvCamera{}
	err := (**C.GError)(C.malloc(C.size_t(unsafe.Sizeof(C.GError{}))))
	out := C.camera_init(cam, err)
	if out != 0 {
		return &Camera{}, fmt.Errorf("camera_init failed with error code %d: %s", out, C.GoString((*err).message))
	}

	return &Camera{ptr: cam}, nil
}

func (cam *Camera) Destroy() {
	C.camera_delete(cam.ptr)
	cam.ptr = nil
}

func (cam *Camera) CaptureBuffer() (*Buffer, error) {
	buf := (**C.ArvBuffer)(C.malloc(C.size_t(unsafe.Sizeof(C.ArvBuffer{}))))
	err := (**C.GError)(C.malloc(C.size_t(unsafe.Sizeof(C.GError{}))))
	out := C.camera_capture_buffer(cam.ptr, buf, err)
	if out != 0 {
		return &Buffer{}, fmt.Errorf("camera_capture_buffer failed with error code %d: %s", out, C.GoString((*err).message))
	}

	return &Buffer{ptr: (*C.ArvBuffer)(*buf)}, nil
}

func (buf *Buffer) Process(filename string) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	C.camera_buffer_process(buf.ptr, cstr)
}

func (buf *Buffer) Clear() {
	C.camera_buffer_clear(buf.ptr)
	buf.ptr = nil
}
