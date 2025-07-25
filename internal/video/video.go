// Il pacchetto video espone le funzioni dell'SDK in C++ di
// ActiveSilicon per la cattura di video.
package video

//#include "stdlib.h"
//#include "video.h"
import "C"
import "unsafe"

func CaptureFrame() string {
	cOutFile := C.captureFrame()
	f := C.GoString(cOutFile)
	C.free(unsafe.Pointer(cOutFile))
	return f
}
