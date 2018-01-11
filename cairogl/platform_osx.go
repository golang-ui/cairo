//+build none

package cairogl

import (
	"log"
	"unsafe"

	"github.com/golang-ui/cairo"
)

type platformSurface struct {
	Buffer  unsafe.Pointer
	Surface *cairo.Surface
}

// EXPEREMENTAL — NOT READY

func newPlatformSurface(width, height int) *platformSurface {
	surface := cairo.QuartzSurfaceCreate(cairo.FormatA1, uint32(width), uint32(height))
	cairo.SurfaceFlush(surface)
	log.Println(surface,
		cairo.ImageSurfaceGetFormat(surface),
		cairo.ImageSurfaceGetWidth(surface),
		cairo.ImageSurfaceGetHeight(surface),
		cairo.ImageSurfaceGetStride(surface),
		cairo.ImageSurfaceGetData(surface),
	)
	return &platformSurface{
		Buffer:  unsafe.Pointer(cairo.ImageSurfaceGetData(surface)),
		Surface: surface,
	}
}

func (p *platformSurface) Destroy() {
	if p.Surface != nil {
		cairo.SurfaceDestroy(p.Surface)
		p.Surface = nil
	}
}
