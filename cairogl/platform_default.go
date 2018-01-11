package cairogl

// #include <stdlib.h>
import "C"
import (
	"unsafe"

	"github.com/golang-ui/cairo"
)

type platformSurface struct {
	Buffer  unsafe.Pointer
	Surface *cairo.Surface
}

func newPlatformSurface(width, height int) *platformSurface {
	const bpp = 4
	buffer := allocSurfaceMemory(width * height * bpp)
	return &platformSurface{
		Buffer: buffer,
		Surface: cairo.ImageSurfaceCreateForData(
			(*byte)(buffer), cairo.FormatArgb32, int32(width), int32(height), int32(bpp*width),
		),
	}
}

func (p *platformSurface) Destroy() {
	if p.Surface != nil {
		cairo.SurfaceDestroy(p.Surface)
		p.Surface = nil
	}
	if p.Buffer != nil {
		C.free(p.Buffer)
		p.Buffer = nil
	}
}

var sizeOfuint8 = unsafe.Sizeof(uint8(0))

func allocSurfaceMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), C.size_t(sizeOfuint8))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}
